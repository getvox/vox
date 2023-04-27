package server

import (
	"context"
	"fmt"
	"time"

	"github.com/iobrother/zoo/core/errors"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/ztimer"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/iobrother/zim/gen/errno"
	"github.com/iobrother/zim/gen/rpc/common"
	"github.com/iobrother/zim/gen/rpc/sess"
	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/conn/internal/client"
	"github.com/iobrother/zim/srv/conn/protocol"
)

const (
	Authing = 1
	Authed  = 2
)

type CmdFunc func(c *Connection, p *protocol.Packet) (err error)

type Server struct {
	opts      Options
	tcpServer *TcpServer
	wsServer  *WsServer
	timer     *ztimer.Timer
	// TODO: 分桶
	connManager *ConnManager
	mapCmdFunc  map[protocol.CmdId]CmdFunc
}

func NewServer(opts ...Option) *Server {
	s := new(Server)
	s.opts = newOptions(opts...)
	s.connManager = NewConnManager()

	if s.opts.TcpAddr != "" {
		s.tcpServer = NewTcpServer(s, s.opts.TcpAddr)
	}

	if s.opts.WsAddr != "" {
		s.wsServer = NewWsServer(s, s.opts.WsAddr)
	}

	s.timer = ztimer.NewTimer(100*time.Millisecond, 20)

	s.registerCmdFunc()

	return s
}

func (s *Server) registerCmdFunc() {
	s.mapCmdFunc = make(map[protocol.CmdId]CmdFunc)
	s.mapCmdFunc[protocol.CmdId_Cmd_Noop] = s.handleNoop
	s.mapCmdFunc[protocol.CmdId_Cmd_Logout] = s.handleLogout
	s.mapCmdFunc[protocol.CmdId_Cmd_Send] = s.handleSend
	s.mapCmdFunc[protocol.CmdId_Cmd_Sync] = s.handleSync
	s.mapCmdFunc[protocol.CmdId_Cmd_MsgAck] = s.handleMsgAck
}

func (s *Server) GetConnManager() *ConnManager {
	return s.connManager
}

func (s *Server) GetServerId() string {
	return s.opts.Id
}

func (s *Server) GetTcpServer() *TcpServer {
	return s.tcpServer
}

func (s *Server) GetWsServer() *WsServer {
	return s.wsServer
}

func (s *Server) GetTimer() *ztimer.Timer {
	return s.timer
}

func (s *Server) Start() error {
	go func() {
		if err := s.consumePush(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		s.timer.Start()
	}()
	go func() {
		if s.tcpServer != nil {
			if err := s.tcpServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()
	go func() {
		if s.wsServer != nil {
			if err := s.wsServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()

	return nil
}

func (s *Server) Stop() error {
	var lastError error
	if s.tcpServer != nil {
		if err := s.tcpServer.Stop(); err != nil {
			lastError = err
		}
	}
	if s.wsServer != nil {
		if err := s.wsServer.Stop(); err != nil {
			lastError = err
		}
	}
	return lastError
}

func (s *Server) consumePush() error {
	// process push message
	pushMsg := new(common.PushMsg)
	topic := fmt.Sprintf("push.online.%s", s.GetServerId())
	nc := runtime.GetNC()
	if _, err := nc.Subscribe(topic, func(msg *nats.Msg) {
		if err := proto.Unmarshal(msg.Data, pushMsg); err != nil {
			log.Errorf("proto.Unmarshal error=(%v)", err)
			return
		}

		log.Infof("receive a message")
		for _, id := range pushMsg.Conns {
			if c := s.GetConnManager().Get(id); c != nil {
				if c.Conn != nil {
					p := protocol.Packet{
						HeaderLen: 20,
						Version:   uint32(c.Version),
						Cmd:       uint32(protocol.CmdId_Cmd_Msg),
						Seq:       0,
						BodyLen:   uint32(len(pushMsg.Msg)),
						Body:      pushMsg.Msg,
					}
					_ = c.WritePacket(&p)
				}
			}
		}
	}); err != nil {
		return err
	}
	return nil
}

func (s *Server) OnOpen(c *Connection) {
	// 10秒钟之内没有认证成功，关闭连接
	c.TimerTask = s.GetTimer().AfterFunc(time.Second*10, func() {
		log.Info("auth timeout...")
		c.Close()
	})
}

func (s *Server) OnClose(c *Connection) {
	log.Infof("client=%s close", c.Uin)

	if c.ID == "" {
		return
	}

	if c != nil {
		req := sess.DisconnectReq{
			Uin:    c.Uin,
			ConnId: c.ID,
		}
		_, _ = client.GetSessClient().Disconnect(context.Background(), &req)
	}

	s.GetConnManager().Remove(c)
}

func (s *Server) OnMessage(p *protocol.Packet, c *Connection) {
	if c.Status == Authing {
		cmd := protocol.CmdId(p.Cmd)
		if cmd != protocol.CmdId_Cmd_Login {
			log.Error("first packet must be cmd_login")
			c.Close()
			return
		}
		if err := s.handleLogin(c, p); err != nil {
			c.Close()
			log.Info("AUTH FAILED")
		} else {
			c.Status = Authed
		}
	} else {
		_ = s.handleProto(c, p)
	}
}

func (s *Server) handleLogin(c *Connection, p *protocol.Packet) (err error) {
	req := &protocol.LoginReq{}
	rsp := &protocol.LoginRsp{}

	defer func() {
		// 不论登录成功与失败，均取消定时任务
		c.TimerTask.Cancel()
		c.TimerTask = nil

		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		log.Error(err)
		return errno.ErrDecodew(errno.WithDetail(err.Error()))
	}

	// TODO: validate
	if req.Uin == "" {
		return errno.ErrValidationw(errno.WithDetail("账号不能为空"))
	}

	log.Infof("handleLogin uin=%s platform=%s token=%s device_id=%s device_name=%s",
		req.Uin, req.Platform, req.Token, req.DeviceId, req.DeviceName)

	reqL := sess.LoginReq{
		Uin:        req.Uin,
		Platform:   req.Platform,
		Server:     s.GetServerId(),
		Token:      req.Token,
		DeviceId:   req.DeviceId,
		DeviceName: req.DeviceName,
		Tag:        req.Tag,
		IsNew:      req.IsNew,
	}
	rspL, err := client.GetSessClient().Login(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	if !req.IsNew && rspL.OtherDeviceId != "" {
		log.Infof("登录冲突 uin=%s cur_device_id=%s conflict_device_id=%s conflict_device_name=%s",
			req.Uin, req.DeviceId, rspL.OtherDeviceId, rspL.OtherDeviceName)
		return errno.ErrLoginConflict()
	}
	// 踢掉旧的连接
	if rspL.OtherDeviceId != "" {
		// TODO: 优化，转发给其他连接服务器
		log.Infof("conflict device id=%s", rspL.OtherDeviceId)
		oldConn := s.GetConnManager().Get(rspL.OtherConnId)
		if oldConn != nil && oldConn.Conn != nil {
			reason := fmt.Sprintf("您的账号在设备%s上登录，如果不是本人操作，您的账号可能被盗", req.DeviceName)
			kick := &protocol.Kick{Reason: reason}
			if b, err := proto.Marshal(kick); err == nil {
				pp := protocol.Packet{
					HeaderLen: p.HeaderLen,
					Version:   p.Version,
					Cmd:       uint32(protocol.CmdId_Cmd_Kick),
					Seq:       0,
					BodyLen:   uint32(len(b)),
					Body:      b,
				}

				oldConn.WritePacket(&pp)
			}
			log.Infof("踢掉客户端 uin=%s conn_id=%s device_id=%s", oldConn.Uin, oldConn.ID, oldConn.DeviceId)
			oldConn.Close()
			s.GetConnManager().Remove(oldConn)
		}
	}

	c.ID = rspL.ConnId
	c.DeviceId = req.DeviceId
	c.Uin = reqL.Uin
	c.Platform = req.Platform
	c.Server = s.GetServerId()
	c.Version = int(p.Version)
	s.GetConnManager().Add(c)

	log.Infof("AUTH SUCC uin=%s", req.Uin)
	return nil
}

func (s *Server) handleLogout(c *Connection, p *protocol.Packet) (err error) {
	log.Infof("client %s logout", c.Uin)
	c.WritePacket(p)
	req := sess.LogoutReq{
		Uin:    c.Uin,
		ConnId: c.ID,
	}
	client.GetSessClient().Logout(context.Background(), &req)
	return
}

func (s *Server) handleProto(c *Connection, p *protocol.Packet) (err error) {
	log.Infof("cmd=%d", p.Cmd)
	cmd := protocol.CmdId(p.Cmd)

	if s.mapCmdFunc[cmd] != nil {
		err = s.mapCmdFunc[cmd](c, p)
	}

	return
}

func (s *Server) handleNoop(c *Connection, p *protocol.Packet) (err error) {
	log.Infof("client %s noop", c.Uin)
	c.WritePacket(p)
	req := sess.HeartbeatReq{
		Uin:    c.Uin,
		ConnId: c.ID,
		Server: c.Server,
	}
	client.GetSessClient().Heartbeat(context.Background(), &req)
	return
}

func (s *Server) responseError(c *Connection, p *protocol.Packet, err error) {
	rsp := &protocol.Error{}
	zerr := errors.FromError(err)
	rsp.Code = zerr.Code
	rsp.Message = zerr.Message
	if zerr.Message == "" {
		rsp.Message = zerr.Detail
	}
	b, _ := proto.Marshal(rsp)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}

func (s *Server) responseMessage(c *Connection, p *protocol.Packet, m proto.Message) {
	b, _ := proto.Marshal(m)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}

package server

import (
	"time"

	"github.com/iobrother/zim/srv/conn/protocol"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/ztimer"
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
	return nil
}

func (s *Server) handleLogout(c *Connection, p *protocol.Packet) (err error) {
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
	return
}

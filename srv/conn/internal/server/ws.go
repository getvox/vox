package server

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/iobrother/zoo/core/log"
	"github.com/panjf2000/gnet/v2"

	"github.com/iobrother/zim/srv/conn/protocol"
)

type WsServer struct {
	gnet.BuiltinEventEngine
	eng   gnet.Engine
	addr  string
	codec *wsCodec
	srv   *Server
}

func NewWsServer(srv *Server, addr string) *WsServer {
	ws := new(WsServer)
	ws.addr = addr
	ws.codec = &wsCodec{}
	ws.srv = srv
	return ws
}

func (ws *WsServer) Start() error {
	return gnet.Run(ws, ws.addr, gnet.WithMulticore(true))
}

func (ws *WsServer) Stop() error {
	return ws.eng.Stop(context.Background())
}

func (ws *WsServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	log.Infof("ws server is listening on %s", ws.addr)
	ws.eng = eng
	return
}

func (ws *WsServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Info("TCP OnOpened ...")
	conn := &Connection{
		Status:      Authing,
		Conn:        c,
		IsWebsocket: true,
	}
	c.SetContext(conn)
	ws.srv.OnOpen(conn)
	return
}

func (ws *WsServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Infof("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}
	ws.srv.OnClose(conn)
	return
}

func (ws *WsServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}

	if ws.codec.readBufferBytes(c) == gnet.Close {
		return gnet.Close
	}
	ok, action = ws.codec.upgrade(c)
	if !ok {
		return
	}
	if ws.codec.buf.Len() <= 0 {
		return gnet.None
	}
	messages, err := ws.codec.Decode(c)
	if err != nil {
		return gnet.Close
	}
	if messages == nil {
		return
	}
	for _, message := range messages {
		p := protocol.Packet{}
		if err := p.Read(message.Payload); err != nil {
			return gnet.Close
		}

		ws.srv.OnMessage(&p, conn)
	}
	return gnet.None
}

type wsCodec struct {
	upgraded bool         // 链接是否升级
	buf      bytes.Buffer // 从实际socket中读取到的数据缓存
	wsMsgBuf wsMessageBuf // ws 消息缓存
}

type wsMessageBuf struct {
	curHeader *ws.Header
	cachedBuf bytes.Buffer
}

type readWrite struct {
	io.Reader
	io.Writer
}

func (w *wsCodec) upgrade(c gnet.Conn) (ok bool, action gnet.Action) {
	if w.upgraded {
		ok = true
		return
	}
	buf := &w.buf
	tmpReader := bytes.NewReader(buf.Bytes())
	oldLen := tmpReader.Len()
	log.Infof("do Upgrade")

	hs, err := ws.Upgrade(readWrite{tmpReader, c})
	skipN := oldLen - tmpReader.Len()
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF { // 数据不完整
			return
		}
		buf.Next(skipN)
		log.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
		action = gnet.Close
		return
	}
	buf.Next(skipN)
	log.Infof("conn[%v] upgrade websocket protocol! Handshake: %v", c.RemoteAddr().String(), hs)
	if err != nil {
		log.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
		action = gnet.Close
		return
	}
	w.upgraded = true
	return true, gnet.None
}

func (w *wsCodec) readBufferBytes(c gnet.Conn) gnet.Action {
	size := c.InboundBuffered()
	buf := make([]byte, size, size)
	read, err := c.Read(buf)
	if err != nil {
		log.Infof("read err! %w", err)
		return gnet.Close
	}
	if read < size {
		log.Infof("read bytes len err! size: %d read: %d", size, read)
		return gnet.Close
	}
	w.buf.Write(buf)
	return gnet.None
}

func (w *wsCodec) Decode(c gnet.Conn) (outs []wsutil.Message, err error) {
	fmt.Println("do Decode")
	messages, err := w.readWsMessages()
	if err != nil {
		log.Infof("Error reading message! %v", err)
		return nil, err
	}
	if messages == nil || len(messages) <= 0 { // 没有读到完整数据 不处理
		return
	}
	for _, message := range messages {
		if message.OpCode.IsControl() {
			err = wsutil.HandleClientControlMessage(c, message)
			if err != nil {
				return
			}
			continue
		}
		if message.OpCode == ws.OpText || message.OpCode == ws.OpBinary {
			outs = append(outs, message)
		}
	}
	return
}

func (w *wsCodec) readWsMessages() (messages []wsutil.Message, err error) {
	msgBuf := &w.wsMsgBuf
	in := &w.buf
	for {
		if msgBuf.curHeader == nil {
			if in.Len() < ws.MinHeaderSize { // 头长度至少是2
				return
			}
			var head ws.Header
			if in.Len() >= ws.MaxHeaderSize {
				head, err = ws.ReadHeader(in)
				if err != nil {
					return messages, err
				}
			} else { // 有可能不完整，构建新的 reader 读取 head 读取成功才实际对 in 进行读操作
				tmpReader := bytes.NewReader(in.Bytes())
				oldLen := tmpReader.Len()
				head, err = ws.ReadHeader(tmpReader)
				skipN := oldLen - tmpReader.Len()
				if err != nil {
					if err == io.EOF || err == io.ErrUnexpectedEOF { // 数据不完整
						return messages, nil
					}
					in.Next(skipN)
					return nil, err
				}
				in.Next(skipN)
			}

			msgBuf.curHeader = &head
			err = ws.WriteHeader(&msgBuf.cachedBuf, head)
			if err != nil {
				return nil, err
			}
		}
		dataLen := (int)(msgBuf.curHeader.Length)
		if dataLen > 0 {
			if in.Len() >= dataLen {
				_, err = io.CopyN(&msgBuf.cachedBuf, in, int64(dataLen))
				if err != nil {
					return
				}
			} else { // 数据不完整
				fmt.Println(in.Len(), dataLen)
				log.Infof("incomplete data")
				return
			}
		}
		if msgBuf.curHeader.Fin { // 当前 header 已经是一个完整消息
			messages, err = wsutil.ReadClientMessage(&msgBuf.cachedBuf, messages)
			if err != nil {
				return nil, err
			}
			msgBuf.cachedBuf.Reset()
		} else {
			log.Infof("The data is split into multiple frames")
		}
		msgBuf.curHeader = nil
	}
}

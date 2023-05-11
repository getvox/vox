package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"

	"github.com/iobrother/zoo/core/log"
	"github.com/panjf2000/gnet/v2"

	"github.com/getvox/vox/srv/conn/protocol"
)

type TcpServer struct {
	gnet.BuiltinEventEngine
	eng   gnet.Engine
	addr  string
	codec *tcpCodec
	srv   *Server
}

func NewTcpServer(srv *Server, addr string) *TcpServer {
	return &TcpServer{
		addr:  addr,
		codec: &tcpCodec{},
		srv:   srv,
	}
}

func (s *TcpServer) Start() error {
	return gnet.Run(s, s.addr, gnet.WithMulticore(true), gnet.WithReusePort(true))
}

func (s *TcpServer) Stop() error {
	return s.eng.Stop(context.Background())
}

func (s *TcpServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	log.Infof("tcp server is listening on %s", s.addr)
	s.eng = eng
	return
}

func (s *TcpServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Info("TCP OnOpened ...")
	conn := &Connection{
		Status: Authing,
		Conn:   c,
	}
	c.SetContext(conn)
	s.srv.OnOpen(conn)
	return
}

func (s *TcpServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Infof("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}
	s.srv.OnClose(conn)
	return
}

func (s *TcpServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	conn, ok := c.Context().(*Connection)
	if !ok {
		return
	}

	for {
		data, err := s.codec.Decode(c)
		if err == ErrIncompletePacket {
			break
		}
		if err != nil {
			log.Errorf("invalid packet: %v", err)
			return gnet.Close
		}

		p := protocol.Packet{}
		if err := p.Read(data); err != nil {
			log.Error(err)
			return gnet.Close
		}

		s.srv.OnMessage(&p, conn)
	}

	return
}

// ==================================== Codec ==============================================

type tcpCodec struct{}

func (_ *tcpCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	return buf, nil
}

var ErrIncompletePacket = errors.New("incomplete packet")

func (_ *tcpCodec) Decode(c gnet.Conn) ([]byte, error) {
	header, _ := c.Peek(protocol.HeaderLen)
	if len(header) < protocol.HeaderLen {
		return nil, ErrIncompletePacket
	}
	byteBuffer := bytes.NewBuffer(header)
	var p protocol.Packet
	if err := binary.Read(byteBuffer, binary.BigEndian, &p.HeaderLen); err != nil {
		return nil, err
	}
	if err := binary.Read(byteBuffer, binary.BigEndian, &p.Version); err != nil {
		return nil, err
	}
	if err := binary.Read(byteBuffer, binary.BigEndian, &p.Cmd); err != nil {
		return nil, err
	}
	if err := binary.Read(byteBuffer, binary.BigEndian, &p.Seq); err != nil {
		return nil, err
	}
	if err := binary.Read(byteBuffer, binary.BigEndian, &p.BodyLen); err != nil {
		return nil, err
	}
	msgLen := int(protocol.HeaderLen + p.BodyLen)
	if c.InboundBuffered() < msgLen {
		return nil, ErrIncompletePacket
	}
	buf, _ := c.Peek(msgLen)
	_, _ = c.Discard(msgLen)
	return buf, nil
}

package server

import (
	"bytes"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/iobrother/ztimer"
	"github.com/panjf2000/gnet/v2"

	"github.com/iobrother/zim/srv/conn/protocol"
)

type Connection struct {
	ID          string
	Status      int
	TimerTask   *ztimer.TimerTask
	DeviceId    string
	Conn        gnet.Conn
	Version     int
	Uin         string
	Platform    string
	Server      string
	IsWebsocket bool
	WsCodec     *wsCodec
}

func (c *Connection) Write(data []byte) error {
	if c.IsWebsocket {
		return wsutil.WriteServerMessage(c.Conn, ws.OpBinary, data)
	} else {
		_, err := c.Conn.Write(data)
		return err
	}
}

func (c *Connection) WritePacket(p *protocol.Packet) error {
	buf := &bytes.Buffer{}
	if err := p.Write(buf); err != nil {
		return err
	}

	if c.IsWebsocket {
		return wsutil.WriteServerMessage(c.Conn, ws.OpBinary, buf.Bytes())
	} else {
		_, err := c.Conn.Write(buf.Bytes())
		return err
	}
}

func (c *Connection) Close() {
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}

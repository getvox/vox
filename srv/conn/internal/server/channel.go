package server

import (
	"context"

	"github.com/iobrother/zoo/core/log"
	"google.golang.org/protobuf/proto"

	"github.com/getvox/vox/gen/rpc/channel"
	"github.com/getvox/vox/srv/conn/internal/client"
	"github.com/getvox/vox/srv/conn/protocol"
)

func (s *Server) handleCreateChannel(c *Connection, p *protocol.Packet) (err error) {
	req := &protocol.CreateChannelReq{}
	rsp := &protocol.CreateChannelRsp{}

	defer func() {
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		return
	}

	reqL := channel.CreateReq{}
	rspL, err := client.GetChannelClient().Create(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}
	rsp.Cid = rspL.Cid

	return nil
}

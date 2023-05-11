package server

import (
	"context"

	"github.com/iobrother/zoo/core/log"
	"google.golang.org/protobuf/proto"

	"github.com/iobrother/zim/gen/rpc/chat"
	"github.com/iobrother/zim/srv/conn/internal/client"
	"github.com/iobrother/zim/srv/conn/protocol"
)

func (s *Server) handleSend(c *Connection, p *protocol.Packet) (err error) {
	log.Info("handleSend ...")
	req := &protocol.SendReq{}
	rsp := &protocol.SendRsp{}

	defer func() {
		if err != nil {
			s.responseError(c, p, err)
		} else {
			s.responseMessage(c, p, rsp)
		}
	}()

	if err = proto.Unmarshal(p.Body, req); err != nil {
		log.Error(err)
		return
	}

	r := chat.SendReq{
		ChannelType:   req.ChannelType,
		MsgType:       req.MsgType,
		From:          req.From,
		To:            req.To,
		Content:       req.Content,
		Uuid:          req.Uuid,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
	}
	rspL, err := client.GetChatClient().SendMsg(context.Background(), &r)
	if err != nil {
		log.Error(err)
		return
	}

	rsp.Id = rspL.Id
	rsp.SendTime = rspL.SendTime
	rsp.Uuid = r.Uuid
	return nil
}

func (s *Server) handleMsgAck(c *Connection, p *protocol.Packet) (err error) {
	req := &protocol.MsgAckReq{}
	rsp := &protocol.MsgAckRsp{}

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

	reqL := chat.MsgAckReq{
		Uin: c.Uin,
		Id:  req.Id,
	}

	_, err = client.GetChatClient().MsgAck(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (s *Server) handleSync(c *Connection, p *protocol.Packet) (err error) {
	req := &protocol.SyncMsgReq{}
	rsp := &protocol.SyncMsgRsp{}

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

	reqL := chat.SyncMsgReq{
		Uin:    c.Uin,
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	rspL, err := client.GetChatClient().SyncMsg(context.Background(), &reqL)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range rspL.List {
		msg := &protocol.Msg{
			Id:          v.Id,
			ChannelType: v.ChannelType,
			Type:        v.Type,
			Content:     v.Content,
			From:        v.From,
			To:          v.To,
			SendTime:    v.SendTime,
			Uuid:        v.Uuid,
			AtUserList:  v.AtUserList,
		}
		rsp.List = append(rsp.List, msg)
	}

	return nil
}

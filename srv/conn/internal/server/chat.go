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
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		Sender:        req.Sender,
		Target:        req.Target,
		Content:       req.Content,
		ClientUuid:    req.ClientUuid,
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
	rsp.ClientUuid = r.ClientUuid
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
			Id:         v.Id,
			ConvType:   v.ConvType,
			Type:       v.Type,
			Content:    v.Content,
			Sender:     v.Sender,
			Target:     v.Target,
			SendTime:   v.SendTime,
			ClientUuid: v.ClientUuid,
			AtUserList: v.AtUserList,
		}
		rsp.List = append(rsp.List, msg)
	}

	return nil
}

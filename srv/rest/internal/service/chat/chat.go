package chat

import (
	"context"
	"sync"

	"github.com/iobrother/zim/gen/http/rest/chat"
	pb "github.com/iobrother/zim/gen/rpc/chat"
	"github.com/iobrother/zim/srv/rest/internal/client"
)

type Service struct{}

var (
	service *Service
	once    sync.Once
)

func GetService() *Service {
	once.Do(func() {
		service = &Service{}
	})
	return service
}

func (s *Service) Send(ctx context.Context, req *chat.SendReq) (rsp *chat.SendRsp, err error) {
	reqL := pb.SendReq{
		ConvType:      req.ConvType,
		MsgType:       req.MsgType,
		From:          req.From,
		To:            req.To,
		Content:       req.Content,
		AtUserList:    nil,
		Uuid:          "",
		IsTransparent: req.IsTransparent,
	}

	cli := client.GetChatClient()
	rspL, err := cli.SendMsg(ctx, &reqL)
	if err != nil {
		return
	}

	rsp.Id = rspL.Id
	rsp.SendTime = rspL.SendTime

	return
}

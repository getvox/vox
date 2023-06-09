package chat

import (
	"context"
	"sync"

	"github.com/getvox/vox/gen/http/rest/chat"
	pb "github.com/getvox/vox/gen/rpc/chat"
	"github.com/getvox/vox/srv/rest/internal/client"
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
		ChannelType:   req.ChannelType,
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

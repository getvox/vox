package channel

import (
	"context"
	"sync"

	"github.com/getvox/vox/gen/http/rest/channel"
	pb "github.com/getvox/vox/gen/rpc/channel"
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

func (s *Service) Create(ctx context.Context, req *channel.CreateReq) (rsp *channel.CreateRsp, err error) {
	reqL := pb.CreateReq{
		Owner:   req.Owner,
		Members: req.Members,
		Name:    req.Name,
		Cid:     req.Cid,
		Notice:  req.Notice,
		Intro:   req.Intro,
		Avatar:  req.Avatar,
	}

	cli := client.GetChannelClient()
	rspL, err := cli.Create(ctx, &reqL)
	if err != nil {
		return
	}
	rsp.Cid = rspL.Cid
	return
}

func (s *Service) Add(ctx context.Context, req *channel.AddReq) (rsp *channel.AddRsp, err error) {
	reqL := pb.InviteUserReq{
		Cid:      req.Cid,
		UserList: req.Members,
	}

	cli := client.GetChannelClient()
	_, err = cli.InviteUser(ctx, &reqL)

	return
}

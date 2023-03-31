package group

import (
	"context"
	"sync"

	"github.com/iobrother/zim/gen/http/rest/group"
	pb "github.com/iobrother/zim/gen/rpc/group"
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

func (s *Service) Create(ctx context.Context, req *group.CreateReq) (rsp *group.CreateRsp, err error) {
	reqL := pb.CreateReq{
		Owner:   req.Owner,
		Members: req.Members,
		Name:    req.Name,
		GroupId: req.GroupId,
		Notice:  req.Notice,
		Intro:   req.Intro,
		Avatar:  req.Avatar,
	}

	cli := client.GetGroupClient()
	rspL, err := cli.Create(ctx, &reqL)
	if err != nil {
		return
	}
	rsp.GroupId = rspL.GroupId
	return
}

func (s *Service) Add(ctx context.Context, req *group.AddReq) (rsp *group.AddRsp, err error) {
	reqL := pb.InviteUserToGroupReq{
		GroupId:  req.GroupId,
		UserList: req.Members,
	}

	cli := client.GetGroupClient()
	_, err = cli.InviteUserToGroup(ctx, &reqL)

	return
}

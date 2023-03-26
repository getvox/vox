package service

import (
	"context"
	"os"
	"sync"

	"github.com/iobrother/zim/srv/gid/internal/snowflake"
	"github.com/spf13/cast"

	"github.com/iobrother/zim/gen/rpc/gid"
)

type Service struct {
	snowflake *snowflake.Snowflake
}

var (
	service *Service
	once    sync.Once
)

func GetService() *Service {
	once.Do(func() {
		service = &Service{}
		s := snowflake.Settings{
			MachineID: getMachineId,
		}
		service.snowflake = snowflake.NewSnowflake(s)
	})
	return service
}

func getMachineId() (uint16, error) {
	var machineId string
	machineId = os.Getenv("MACHINE_ID")
	if machineId == "" {
		return 1, nil
	}

	return cast.ToUint16(machineId), nil
}

func (s *Service) Get(ctx context.Context, req *gid.GetReq, rsp *gid.GetRsp) (err error) {
	rsp.Id = s.snowflake.NextId()
	return
}

func (s *Service) GetBatch(ctx context.Context, req *gid.GetBatchReq, rsp *gid.GetBatchRsp) (err error) {
	for i := 0; i < int(req.Count); i++ {
		rsp.Ids = append(rsp.Ids, s.snowflake.NextId())
	}
	return
}

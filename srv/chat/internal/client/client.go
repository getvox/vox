package client

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/zoo/core/transport/rpc/client"

	"github.com/getvox/vox/gen/rpc/gid"
)

var gidClient *gid.GidClient

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetGidClient() *gid.GidClient {
	if gidClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("GetGidClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Gid"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("GetGidClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		gidClient = gid.NewGidClient(cli)
	}
	return gidClient
}

package client

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/zoo/core/transport/rpc/client"

	"github.com/getvox/vox/gen/rpc/sess"
)

var sessClient *sess.SessClient

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetSessClient() *sess.SessClient {
	if sessClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("getSessClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Sess"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("getSessClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		sessClient = sess.NewSessClient(cli)
	}
	return sessClient
}

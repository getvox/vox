package main

import (
	"github.com/iobrother/zoo"
	"github.com/iobrother/zoo/core/log"
	"github.com/smallnest/rpcx/server"

	"github.com/getvox/vox/pkg/runtime"
	"github.com/getvox/vox/srv/sess/internal/service"
)

func main() {
	app := zoo.New(
		zoo.InitRpcServer(InitRpcServer),
		zoo.BeforeStart(func() error {
			runtime.Setup()
			return nil
		}),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Sess", service.GetService(), ""); err != nil {
		log.Fatal(err)
	}
	return nil
}

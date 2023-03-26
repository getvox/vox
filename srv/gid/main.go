package main

import (
	"github.com/iobrother/zoo"
	"github.com/iobrother/zoo/core/log"
	"github.com/smallnest/rpcx/server"

	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/gid/internal/service"
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
	if err := s.RegisterName("Gid", service.GetService(), ""); err != nil {
		log.Fatal(err)
	}
	return nil
}

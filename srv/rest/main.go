package main

import (
	"github.com/iobrother/zoo"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/zoo/core/transport/http/server"

	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/rest/internal/router"
)

func main() {
	app := zoo.New(
		zoo.InitHttpServer(InitHttpServer),
		zoo.BeforeStart(func() error {
			runtime.Setup()
			return nil
		}),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitHttpServer(r *server.Server) error {
	router.Setup(r)
	return nil
}

package main

import (
	"github.com/iobrother/zoo"
	"github.com/iobrother/zoo/core/log"
	"github.com/smallnest/rpcx/server"

	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/chat/internal/model"
	"github.com/iobrother/zim/srv/chat/internal/service"
)

func main() {
	app := zoo.New(
		zoo.InitRpcServer(InitRpcServer),
		zoo.BeforeStart(before),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Chat", service.GetChatService(), ""); err != nil {
		log.Fatal(err)
	}
	return nil
}

func before() error {
	runtime.Setup()
	db := runtime.GetDB()
	if err := db.AutoMigrate(
		&model.Msg{},
		&model.User{},
		&model.Channel{},
		&model.Member{},
	); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

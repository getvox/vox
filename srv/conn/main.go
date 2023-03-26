package main

import (
	"github.com/iobrother/zoo/core/log"

	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/conn/internal/app"
)

func main() {
	a := app.New(app.BeforeStart(func() error {
		runtime.Setup()
		return nil
	}))

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

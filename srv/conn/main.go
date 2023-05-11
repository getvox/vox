package main

import (
	"github.com/iobrother/zoo/core/log"

	"github.com/getvox/vox/pkg/runtime"
	"github.com/getvox/vox/srv/conn/internal/app"
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

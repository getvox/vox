package main

import (
	"errors"

	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/srv/task/internal/app"
	"github.com/iobrother/zim/srv/task/internal/model"
	"github.com/iobrother/zoo/core/log"
	"github.com/nats-io/nats.go"
)

func main() {
	a := app.New(app.BeforeStart(before))

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

func before() error {
	runtime.Setup()
	js := runtime.GetJS()
	if _, err := js.StreamInfo("MSGS"); err != nil {
		if !errors.Is(err, nats.ErrStreamNotFound) {
			return err
		}
		// nats stream add MSGS --subjects "MSGS.*" --ack --max-msgs=-1 --max-bytes=-1 --max-age=-1 --storage file --retention work --max-msg-size=-1 --discard=old
		js.AddStream(&nats.StreamConfig{
			Name:      "MSGS",
			Subjects:  []string{"MSGS.*"},
			Retention: nats.WorkQueuePolicy,
			Storage:   nats.FileStorage,
		})
	}

	if _, err := js.ConsumerInfo("MSGS", "TASK_NEW"); err != nil {
		if !errors.Is(err, nats.ErrConsumerNotFound) {
			return err
		}
		// nats consumer add MSGS TASK.new --filter MSGS.received --ack explicit --pull --deliver all --max-deliver=-1
		if _, err := js.AddConsumer("MSGS", &nats.ConsumerConfig{
			Durable:       "TASK_NEW",
			AckPolicy:     nats.AckExplicitPolicy,
			FilterSubject: "MSGS.new",
		}); err != nil {
			log.Error(err)
			return err
		}
	}

	if _, err := js.ConsumerInfo("MSGS", "TASK_TODO"); err != nil {
		if !errors.Is(err, nats.ErrConsumerNotFound) {
			return err
		}
		// nats consumer add MSGS TASK.todo --filter MSGS.todo --ack explicit --pull --deliver all --max-deliver=-1
		if _, err := js.AddConsumer("MSGS", &nats.ConsumerConfig{
			Durable:       "TASK_TODO",
			AckPolicy:     nats.AckExplicitPolicy,
			FilterSubject: "MSGS.todo",
		}); err != nil {
			log.Error(err)
			return err
		}
	}

	db := runtime.GetDB()
	if err := db.AutoMigrate(&model.Msg{}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

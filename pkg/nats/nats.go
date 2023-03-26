package nats

import "github.com/nats-io/nats.go"

type Config struct {
	Addr string
}

func Open(c *Config) (*nats.Conn, error) {
	nc, err := nats.Connect(c.Addr)
	if err != nil {
		return nil, err
	}

	return nc, nil
}

package server

import "github.com/google/uuid"

type Options struct {
	Id      string
	TcpAddr string
	WsAddr  string
}

func newOptions(opts ...Option) Options {
	options := Options{
		Id: uuid.New().String(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func TcpAddr(addr string) Option {
	return func(o *Options) {
		o.TcpAddr = addr
	}
}

func WsAddr(addr string) Option {
	return func(o *Options) {
		o.WsAddr = addr
	}
}

package app

type Options struct {
	BeforeStart []func() error
}

func newOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func BeforeStart(f func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, f)
	}
}

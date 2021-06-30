package Options

import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeOut = 10
	defaultCaching = true
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(o *options)
}

type optionFuc func(*options)

func (f optionFuc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFuc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(b bool) Option {
	return optionFuc(func(o *options) {
		o.caching = b
	})
}

type ConnectBuilder struct {
	opts []Option
}

func (b ConnectBuilder) WithCache(cache bool) ConnectBuilder {
	b.opts = append(b.opts, WithCaching(cache))
	return b
}

func (b ConnectBuilder) WithTimeOut(t time.Duration) ConnectBuilder {
	b.opts = append(b.opts, WithTimeout(t))
	return b
}

func Connect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeOut,
		caching: defaultCaching,
	}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

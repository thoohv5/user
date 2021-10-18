package base

import (
	"context"

	"github.com/thoohv5/template/internal/ent"
	"github.com/thoohv5/template/pkg/log"
)

type Base struct {
	options
}

func New(opts ...Option) *Base {
	options := options{}
	for _, o := range opts {
		o.apply(&options)
	}
	return &Base{
		options: options,
	}
}

type options struct {
	tx     *ent.Tx
	log    log.ILog
	client *ent.Client
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(opts *options) {
	f(opts)
}

func WithTx(tx *ent.Tx) Option {
	return optionFunc(func(o *options) {
		o.tx = tx
	})
}

func WithLog(log log.ILog) Option {
	return optionFunc(func(o *options) {
		o.log = log
	})
}

func WithClient(client *ent.Client) Option {
	return optionFunc(func(o *options) {
		o.client = client
	})
}

// GetTx Tx
func (b *Base) GetTx() *ent.Tx {
	return b.options.tx
}

// GetLogger log
func (b *Base) GetLogger() log.ILog {
	return b.options.log
}

// GetClient Client
func (b *Base) GetClient() *ent.Client {
	return b.options.client
}

// Exist 是否存在
func (b *Base) Exist(ctx context.Context, identity string) (bool, error) {
	return true, nil
}

// Create 创建
func (b *Base) Create(ctx context.Context, identity string, extra string) error {
	return nil
}

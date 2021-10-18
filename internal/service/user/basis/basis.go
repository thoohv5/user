// Package basis 用户基础
package basis

import (
	"context"

	"go.uber.org/zap"

	"github.com/thoohv5/template/internal/codes"
	"github.com/thoohv5/template/internal/ent"
	"github.com/thoohv5/template/pkg/log"
)

type basis struct {
	options
}

// IBasis 用户基础标准
type IBasis interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, userIdentity string, extra string) error
	// CreateUserInfo 创建用户信息
	CreateUserInfo(ctx context.Context, userIdentity string, extra string) error
	// CreateUserExtend 创建用户扩展
	CreateUserExtend(ctx context.Context, userIdentity string, extra string) error

	// UpdateUser 修改用户
	UpdateUser(ctx context.Context, userIdentity string, extra string) error
	// UpdateUserInfo 修改用户信息
	UpdateUserInfo(ctx context.Context, userIdentity string, extra string) error
	// UpdateUserExtend 修改用户扩展
	UpdateUserExtend(ctx context.Context, userIdentity string, extra string) error
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

func WithLog(log log.ILog) Option {
	return optionFunc(func(o *options) {
		o.log = log
	})
}

func WithTx(tx *ent.Tx) Option {
	return optionFunc(func(o *options) {
		o.tx = tx
	})
}

func New(opts ...Option) IBasis {
	options := options{}
	for _, o := range opts {
		o.apply(&options)
	}
	return &basis{
		options,
	}
}

// CreateUser 创建用户
func (b *basis) CreateUser(ctx context.Context, userIdentity string, extra string) error {
	// 创建用户
	_, err := b.tx.User.Create().SetIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("User Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

// CreateUserInfo 创建用户信息
func (b *basis) CreateUserInfo(ctx context.Context, userIdentity string, extra string) error {
	// 创建用户信息
	_, err := b.tx.UserInfo.Create().SetUserIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("UserInfo Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

// CreateUserExtend 创建用户扩展
func (b *basis) CreateUserExtend(ctx context.Context, userIdentity string, extra string) error {
	// 创建用户扩展信息
	_, err := b.tx.UserExtend.Create().SetUserIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("UserExtend Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

// UpdateUser 修改用户
func (b *basis) UpdateUser(ctx context.Context, userIdentity string, extra string) error {
	// 修改用户
	_, err := b.tx.User.Update().SetIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("User Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

// UpdateUserInfo 修改用户信息
func (b *basis) UpdateUserInfo(ctx context.Context, userIdentity string, extra string) error {
	// 修改用户信息
	_, err := b.tx.UserInfo.Update().SetUserIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("UserInfo Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

// UpdateUserExtend 修改用户扩展
func (b *basis) UpdateUserExtend(ctx context.Context, userIdentity string, extra string) error {
	// 修改用户扩展信息
	_, err := b.tx.UserExtend.Update().SetUserIdentity(userIdentity).Save(ctx)
	if nil != err {
		b.log.Error("UserExtend Save err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return codes.ErrDbException
	}
	return nil
}

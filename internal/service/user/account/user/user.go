package user

import (
	"context"
	"strconv"

	"go.uber.org/zap"

	"github.com/thoohv5/template/internal/codes"
	"github.com/thoohv5/template/internal/ent/useraccount"
	"github.com/thoohv5/template/internal/service/user/account/base"
	"github.com/thoohv5/template/internal/service/user/account/standard"
	"github.com/thoohv5/template/internal/service/user/entity"
)

type user struct {
	*base.Base
}

func New(opts ...base.Option) standard.IAccount {
	return &user{
		Base: base.New(opts...),
	}
}

// Exist 是否存在
func (u *user) Exist(ctx context.Context, identity string) (bool, error) {

	// 类型转换
	account, err := strconv.ParseInt(identity, 10, 32)
	if nil != err {
		u.GetLogger().Error("strconv ParseInt err", zap.Error(err), zap.String("identity", identity))
		return false, codes.ErrTypeConvertException
	}

	return u.GetTx().UserAccount.Query().Where(useraccount.Account(account)).Exist(ctx)
}

// Create 创建
func (u *user) Create(ctx context.Context, userIdentity string, extra string) error {

	_, err := u.GetTx().UserAccount.Create().SetUserIdentity(userIdentity).Save(ctx)
	if nil != err {
		u.GetLogger().Error("UserAccount Save err", zap.Error(err), zap.String("userIdentity", userIdentity), zap.String("extra", extra))
		return codes.ErrDbException
	}

	return nil
}

// Detail 详情
func (u *user) Detail(ctx context.Context, identity string, resp *entity.DetailEntity) (userIdentity string, err error) {

	userAccount, err := u.GetClient().UserAccount.Query().Where(useraccount.UserIdentity(identity)).First(ctx)
	if nil != err {
		u.GetLogger().Error("UserAccount First err", zap.Error(err), zap.String("identity", identity))
		return "", codes.ErrDbException
	}

	resp.UserAccount = userAccount

	return userAccount.UserIdentity, nil

}

// Update 详情
func (u *user) Update(ctx context.Context, userIdentity string, extra string) error {
	return nil
}

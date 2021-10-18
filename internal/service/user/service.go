// Package user 业务
package user

import (
	"context"

	"go.uber.org/zap"

	"github.com/thoohv5/template/internal/codes"
	"github.com/thoohv5/template/internal/data"
	"github.com/thoohv5/template/internal/ent"
	"github.com/thoohv5/template/internal/ent/user"
	"github.com/thoohv5/template/internal/ent/userextend"
	"github.com/thoohv5/template/internal/ent/userinfo"
	"github.com/thoohv5/template/internal/enum"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/pkg/dbx"
	"github.com/thoohv5/template/internal/service/user/account"
	"github.com/thoohv5/template/internal/service/user/account/base"
	"github.com/thoohv5/template/internal/service/user/basis"
	"github.com/thoohv5/template/internal/service/user/entity"
	"github.com/thoohv5/template/pkg/log"
)

type service struct {
	// 配置
	cf config.IConfig
	// 日志
	log log.ILog

	// 数据源
	data data.IData
}

// New 创建
func New(cf config.IConfig, log log.ILog, data data.IData) IService {
	return &service{
		cf:  cf,
		log: log,

		data: data,
	}
}

// Create 创建
func (s *service) Create(ctx context.Context, accountType enum.AccountType, identity string, extra string) (userIdentity string, err error) {

	userIdentity = ""

	tx, err := s.data.GetEdb().Primary().Tx(ctx)
	if nil != err {
		s.log.Error("db get tx err", zap.Error(err))
		return userIdentity, codes.ErrDbException
	}

	// 适配器
	accountAdapter := account.Adapter(accountType, base.WithLog(s.log), base.WithTx(tx))
	bb := basis.New(basis.WithLog(s.log), basis.WithTx(tx))

	// 创建用户
	if err := dbx.ExecTrans(tx, func(tx *ent.Tx) error {
		// 检查账户是否存在
		exist, err := accountAdapter.Exist(ctx, identity)
		if nil != err {
			s.log.Error("accountAdapter Exist err", zap.Error(err), zap.String("userIdentity", identity), zap.Int32("accountType", accountType))
			return err
		}
		// 存在
		if exist {
			s.log.Error("Having Exist err", zap.String("userIdentity", identity), zap.Int32("accountType", accountType))
			return codes.ErrHavingExist
		}
		return nil
	}, func(tx *ent.Tx) error {
		// 创建用户
		return bb.CreateUser(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 创建用户信息
		return bb.CreateUserInfo(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 创建用户扩展信息
		return bb.CreateUserExtend(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 用户账户创建
		if err := accountAdapter.Create(ctx, userIdentity, extra); nil != err {
			s.log.Error("accountAdapter Create err", zap.Error(err), zap.String("userIdentity", userIdentity), zap.String("extra", extra))
			return err
		}
		return nil
	}); nil != err {
		s.log.Error("user ExecTrans err", zap.Error(err), zap.String("userIdentity", identity), zap.Int32("accountType", accountType), zap.String("extra", extra))
		return userIdentity, codes.ErrDbTransactionException
	}

	return userIdentity, nil

}

// Detail 详情
func (s *service) Detail(ctx context.Context, accountType enum.AccountType, identity string) (*entity.DetailEntity, error) {
	resp := new(entity.DetailEntity)

	edb := s.data.GetEdb().Replica()

	// 账户信息
	userIdentity, err := account.Adapter(accountType, base.WithClient(edb)).Detail(ctx, identity, resp)
	if nil != err {
		s.log.Error("adapter Detail err", zap.Error(err), zap.String("userIdentity", identity), zap.Int32("accountType", accountType))
		return resp, err
	}

	// 用户
	userEntity, err := edb.User.Query().Where(user.Identity(userIdentity)).First(ctx)
	if nil != err {
		s.log.Error("User First err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return resp, err
	}
	resp.UserEntity = userEntity

	// 用户信息
	userInfoEntity, err := edb.UserInfo.Query().Where(userinfo.UserIdentity(userIdentity)).First(ctx)
	if nil != err {
		s.log.Error("UserInfo First err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return resp, err
	}
	resp.UserInfoEntity = userInfoEntity

	// 用户扩展信息
	userExtendEntity, err := edb.UserExtend.Query().Where(userextend.UserIdentity(userIdentity)).First(ctx)
	if nil != err {
		s.log.Error("UserInfo First err", zap.Error(err), zap.String("userIdentity", userIdentity))
		return resp, err
	}
	resp.UserExtendEntity = userExtendEntity

	return resp, nil
}

// List 列表
func (s *service) List(ctx context.Context, req *entity.ListParam) ([]*entity.UserBasis, error) {

	list := make([]*entity.UserBasis, 0)

	edb := s.data.GetEdb().Replica()

	// 用户列表
	userList, err := edb.User.Query().Where(user.Type(req.Type)).Offset(int(req.BasePage.Start)).Limit(int(req.BasePage.Limit)).All(ctx)
	if nil != err {
		s.log.Error("User All err", zap.Error(err), zap.Any("req", req))
		return list, err
	}

	userIdentityList := make([]string, 0)
	for _, item := range userList {
		userIdentityList = append(userIdentityList, item.Identity)
	}

	// 用户信息
	userInfoByIdentity := make(map[string]*ent.UserInfo, 0)
	userInfoList, err := edb.UserInfo.Query().Where(userinfo.UserIdentityIn(userIdentityList...)).All(ctx)
	if nil != err {
		s.log.Error("UserInfo All err", zap.Error(err), zap.Strings("userIdentityList", userIdentityList))
		return list, err
	}
	for _, item := range userInfoList {
		userInfoByIdentity[item.UserIdentity] = item
	}

	// 用户扩展信息
	userExtendByIdentity := make(map[string]*ent.UserExtend, 0)
	userExtendList, err := edb.UserExtend.Query().Where(userextend.UserIdentityIn(userIdentityList...)).All(ctx)
	if nil != err {
		s.log.Error("UserInfo All err", zap.Error(err), zap.Strings("userIdentityList", userIdentityList))
		return list, err
	}
	for _, item := range userExtendList {
		userExtendByIdentity[item.UserIdentity] = item
	}

	for _, item := range userList {
		identity := item.Identity

		userInfo, ok := userInfoByIdentity[identity]
		if !ok {
			userInfo = new(ent.UserInfo)
		}

		userExtend, ok := userExtendByIdentity[identity]
		if !ok {
			userExtend = new(ent.UserExtend)
		}

		list = append(list, &entity.UserBasis{
			UserEntity:       item,
			UserInfoEntity:   userInfo,
			UserExtendEntity: userExtend,
		})
	}

	return list, nil

}

// Edit 编辑
func (s *service) Edit(ctx context.Context, userIdentity string, param *entity.EditParam) error {

	accountType := param.AccountType
	extra := param.Extra

	tx, err := s.data.GetEdb().Primary().Tx(ctx)
	if nil != err {
		s.log.Error("db get tx err", zap.Error(err))
		return codes.ErrDbException
	}

	// 适配器
	accountAdapter := account.Adapter(accountType, base.WithLog(s.log), base.WithTx(tx))
	bb := basis.New(basis.WithLog(s.log), basis.WithTx(tx))

	// 创建用户
	if err := dbx.ExecTrans(tx, func(tx *ent.Tx) error {
		// 用户查找
		_, err := tx.User.Query().Where(user.Identity(userIdentity)).First(ctx)
		if nil != err {
			s.log.Error("User First err", zap.Error(err), zap.String("userIdentity", userIdentity))
			return codes.ErrDbException
		}
		return nil
	}, func(tx *ent.Tx) error {
		// 修改用户
		return bb.UpdateUser(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 修改用户信息
		return bb.UpdateUserInfo(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 修改用户扩展信息
		return bb.UpdateUserExtend(ctx, userIdentity, extra)
	}, func(tx *ent.Tx) error {
		// 用户账户修改
		if err := accountAdapter.Update(ctx, userIdentity, extra); nil != err {
			s.log.Error("accountAdapter Update err", zap.Error(err), zap.String("userIdentity", userIdentity), zap.String("extra", extra))
			return err
		}
		return nil
	}); nil != err {
		s.log.Error("user ExecTrans err", zap.Error(err), zap.String("userIdentity", userIdentity), zap.Any("req", param))
		return codes.ErrDbTransactionException
	}

	return nil

}

package entity

import (
	"github.com/thoohv5/template/internal/ent"
	"github.com/thoohv5/template/internal/enum"
	"github.com/thoohv5/template/pkg/http"
)

// DetailEntity 详情 返回
type DetailEntity struct {
	UserBasis
	Account
}

// ListParam 列表 请求
type ListParam struct {
	// 分页
	http.BasePage
	// 用户类型
	Type int32
}

// EditParam 编辑 请求
type EditParam struct {
	// 账号类型：0-用户密码，1-手机号码
	AccountType enum.AccountType
	// 唯一标识
	Identity string
	// 额外字段，字段为json，格式：
	Extra string
}

// UserBasis Entity
type UserBasis struct {
	UserEntity       *ent.User
	UserInfoEntity   *ent.UserInfo
	UserExtendEntity *ent.UserExtend
}

// Account Entity
type Account struct {
	UserAccount *ent.UserAccount
}

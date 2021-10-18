package user

import (
	"github.com/thoohv5/template/internal/enum"
	"github.com/thoohv5/template/pkg/http"
)

// RegisterReq 注册 请求
type RegisterReq struct {
	// 账号类型：0-用户密码，1-手机号码
	AccountType enum.AccountType `form:"accountType" json:"accountType"`
	// 唯一标识
	Identity string `form:"identity" json:"identity" binding:"required,gte=1"`
	// 额外字段，字段为json，格式：
	Extra string `form:"extra" json:"extra" binding:"isdefault|json"`
}

// DetailReq 详情 请求
type DetailReq struct {
	// 账号类型：0-用户密码，1-手机号码
	AccountType enum.AccountType `form:"accountType" json:"accountType"`
	// 唯一标识
	Identity string `form:"identity" json:"identity" binding:"omitempty,gte=1"`
}

// ListReq 列表 请求
type ListReq struct {
	// 分页
	http.BasePage
	// 用户类型
	Type int32 `form:"userType" json:"userType" binding:"omitempty,gte=1"`
}

// EditReq 编辑 请求
type EditReq struct {
	// 用户标识
	UserIdentity string `form:"userIdentity" json:"userIdentity" binding:"required,gte=1"`
	// 账号类型：0-用户密码，1-手机号码
	AccountType enum.AccountType `form:"accountType" json:"accountType"`
	// 唯一标识
	Identity string `form:"identity" json:"identity" binding:"omitempty,gte=1"`
	// 额外字段，字段为json，格式：
	Extra string `form:"extra" json:"extra" binding:"isdefault|json"`
}

package user

import "github.com/thoohv5/template/internal/service/user/entity"

// ListResp 列表 返回
type ListResp struct {
	List []*entity.UserBasis
}

// EditResp 编辑 返回
type EditResp struct {
	// 用户标识
	UserIdentity string
}

// RegisterResp 创建 返回
type RegisterResp struct {
	// 用户标识
	UserIdentity string
}

// DetailResp 详情 返回
type DetailResp struct {
	Detail *entity.DetailEntity
}

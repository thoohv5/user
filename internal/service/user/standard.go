package user

import (
	"context"

	"github.com/thoohv5/template/internal/enum"
	"github.com/thoohv5/template/internal/service/user/entity"
)

// IService 业务标准
type IService interface {
	// Create 创建
	Create(ctx context.Context, accountType enum.AccountType, identity string, extra string) (userIdentity string, err error)
	// Detail 详情
	Detail(ctx context.Context, accountType enum.AccountType, identity string) (*entity.DetailEntity, error)
	// List 列表
	List(ctx context.Context, param *entity.ListParam) ([]*entity.UserBasis, error)
	// Edit 编辑
	Edit(ctx context.Context, userIdentity string, param *entity.EditParam) error
}

package standard

import (
	"context"

	"github.com/thoohv5/template/internal/service/user/entity"
)

// IAccount 账号标准
type IAccount interface {
	// Exist 是否存在
	Exist(ctx context.Context, identity string) (bool, error)
	// Create 创建
	Create(ctx context.Context, userIdentity string, extra string) error
	// Update 修改
	Update(ctx context.Context, userIdentity string, extra string) error
	// Detail 详情
	Detail(ctx context.Context, identity string, resp *entity.DetailEntity) (userIdentity string, err error)
}

package v1

import (
	"context"
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
	"ent-three-layer/app/user/srv/internal/domain/do"
	v1 "ent-three-layer/pkg/common/meta/v1"
)

type UserStore interface {
	/*
		有数据访问的方法，一定要有error
		参数中最好有ctx
	*/
	// List 用户列表
	List(ctx context.Context, orderby []string, opts v1.ListMeta) (*do.UserDOList, error)

	// GetByMobile 通过手机号码查询用户
	GetByMobile(ctx context.Context, mobile string) (*ent.User, error)

	// GetByID 通过用户ID查询用户
	GetByID(ctx context.Context, id uint64) (*ent.User, error)

	// Create 创建用户
	Create(ctx context.Context, user *ent.User) error

	// Update 更新用户
	Update(ctx context.Context, user *ent.User) error
}

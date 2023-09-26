package db

import (
	"context"
	v12 "ent-three-layer/app/user/srv/internal/data/v1"
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
	"ent-three-layer/app/user/srv/internal/domain/do"
	v1 "ent-three-layer/pkg/common/meta/v1"
)

type user struct {
	db *ent.Client
}

func newUser(factory *mysqlFactory) *user {
	user := &user{
		db: factory.db,
	}
	return user
}

func (u *user) List(ctx context.Context, orderby []string, opts v1.ListMeta) (*do.UserDOList, error) {
	ret := &do.UserDOList{}

	//分页
	var limit, offset int
	if opts.PageSize == 0 {
		limit = 10
	} else {
		limit = opts.PageSize
	}

	if opts.Page > 0 {
		offset = (opts.Page - 1) * limit
	}
	// 排序
	user, _ := u.db.User.Query().Order(ent.Desc(orderby...)).Offset(offset).Limit(limit).All(ctx)
	count := u.db.User.Query().Aggregate(ent.Count()).Scan(ctx, &ret)

	return &do.UserDOList{Items: rv, TotalCount: int64(count)}, err
}

func (u *user) GetByMobile(ctx context.Context, mobile string) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) GetByID(ctx context.Context, id uint64) (*ent.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) Create(ctx context.Context, user *ent.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) Update(ctx context.Context, user *ent.User) error {
	//TODO implement me
	panic("implement me")
}

var _ v12.UserStore = (*user)(nil)

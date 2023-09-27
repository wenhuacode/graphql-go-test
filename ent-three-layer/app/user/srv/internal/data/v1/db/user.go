package db

import (
	"context"
	v12 "ent-three-layer/app/user/srv/internal/data/v1"
	"ent-three-layer/app/user/srv/internal/data/v1/ent"
	user2 "ent-three-layer/app/user/srv/internal/data/v1/ent/user"
	"ent-three-layer/app/user/srv/internal/domain/do"
	code2 "ent-three-layer/gmicro/code"
	v1 "ent-three-layer/pkg/common/meta/v1"
	"ent-three-layer/pkg/errors"
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

func (u *user) List(ctx context.Context, orderby []v1.OrderMeta, opts v1.ListMeta) (*do.UserDOList, error) {
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

	var order []user2.OrderOption
	// 排序
	for _, value := range orderby {
		if value.Order == 0 {
			order = append(order, ent.Asc(value.Field))
		}
		if value.Order == 1 {
			order = append(order, ent.Desc(value.Field))
		}
	}

	err := u.db.User.Query().Limit(limit).Offset(offset).Order(order...).Aggregate().Scan(ctx, &ret.Items)
	i, _ := u.db.User.Query().Count(ctx)
	ret.TotalCount = int64(i)

	if err != nil {
		return nil, errors.WithCode(code2.ErrDatabase, err.Error())
	}
	return ret, nil
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

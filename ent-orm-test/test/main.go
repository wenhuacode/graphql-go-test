package main

import (
	"context"
	start "ent-orm-test"
	"ent-orm-test/ent"
	user2 "ent-orm-test/ent/user"
	"fmt"
	"github.com/rs/xid"
	"log"
	"strconv"
)

type UserDO struct {
	ent.User
}

type UserDOList struct {
	TotalCount int64     `json:"totalCount,omitempty"`
	Items      []*UserDO `json:"items"`
}

type ListMeta struct {
	Page int `json:"totalCount,omitempty"`

	PageSize int `json:"offset,omitempty" form:"offset"`
}

type OrderMeta struct {
	Field string

	Order int
}

func CreateUser(ctx context.Context, client *ent.Client) {
	save, _ := client.User.Create().
		SetNickname("wenhua").
		SetMobile("13888888888").
		SetPassword("123456").
		SetGender("male").
		SetRole(1).
		Save(ctx)
	fmt.Println(save)
}

func DeleteUser(ctx context.Context, client *ent.Client) {
	fromString, _ := xid.FromString("ck9s2ubskcpivj339e10")
	_ = client.User.DeleteOneID(fromString).Exec(ctx)
}

func getUser(ctx context.Context, client *ent.Client) {
	fromString, _ := xid.FromString("ck9s2ubskcpivj339e10")
	user, _ := client.User.Query().Where(user2.ID(fromString)).First(ctx)
	fmt.Println(user)
}

func mapCreateUser(ctx context.Context, client *ent.Client) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		_, _ = client.User.Create().
			SetMobile("1388888888" + strconv.Itoa(i)).
			SetPassword("12345").
			SetNickname("wenhua").
			SetGender("male").SetRole(0).Save(ctx)
	}
}

func batchQueryUser(ctx context.Context, orderby []OrderMeta, opts ListMeta, client *ent.Client) *UserDOList {
	ret := &UserDOList{}

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

	err := client.User.Query().Limit(limit).Offset(offset).Order(order...).Aggregate().Scan(ctx, &ret.Items)
	i, _ := client.User.Query().Count(ctx)
	ret.TotalCount = int64(i)

	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	return ret
}

func main() {
	client, err := start.Open()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	//DeleteUser(ctx, client)
	//getUser(ctx, client)
	//mapCreateUser(ctx, client)

	var order []OrderMeta
	order = append(order, OrderMeta{
		Field: "id",
		Order: 0,
	})

	var opts ListMeta
	opts.Page = 1
	opts.PageSize = 10
	batchQueryUser(ctx, order, opts, client)
}

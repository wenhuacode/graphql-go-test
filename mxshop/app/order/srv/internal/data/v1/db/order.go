package db

import (
	"context"
	code2 "mxshop/gmicro/code"
	"mxshop/pkg/errors"

	"gorm.io/gorm"
	v1 "mxshop/app/order/srv/internal/data/v1"
	"mxshop/app/order/srv/internal/domain/do"
	metav1 "mxshop/pkg/common/meta/v1"
)

type orders struct {
	db *gorm.DB
}

func newOrders(factory *dataFactory) *orders {
	return &orders{
		db: factory.db,
	}
}

func (o *orders) Get(ctx context.Context, orderSn string) (*do.OrderInfoDO, error) {
	var order do.OrderInfoDO
	err := o.db.WithContext(ctx).Preload("OrderGoods").Where("order_sn = ?", orderSn).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *orders) List(ctx context.Context, userID uint64, meta metav1.ListMeta, orderby []string) (*do.OrderInfoDOList, error) {
	ret := &do.OrderInfoDOList{}
	//分页
	var limit, offset int
	if meta.PageSize == 0 {
		limit = 10
	} else {
		limit = meta.PageSize
	}

	if meta.Page > 0 {
		offset = (meta.Page - 1) * limit
	}

	//排序
	query := o.db.Preload("OrderGoods")
	for _, value := range orderby {
		query = query.Order(value)
	}

	d := query.Offset(offset).Limit(limit).Find(&ret.Items).Count(&ret.TotalCount)
	if d.Error != nil {
		return nil, errors.WithCode(code2.ErrDatabase, d.Error.Error())
	}
	return ret, nil
}

// Create 创建订单之后要删除对应的购物车记录
func (o *orders) Create(ctx context.Context, txn *gorm.DB, order *do.OrderInfoDO) error {
	db := o.db
	if txn != nil {
		db = txn
	}
	return db.Create(order).Error
}

func (o *orders) Update(ctx context.Context, txn *gorm.DB, order *do.OrderInfoDO) error {
	db := o.db
	if txn != nil {
		db = txn
	}
	return db.Model(order).Save(order).Error
}

var _ v1.OrderStore = &orders{}

package main

import (
	"context"
	"fmt"
)

type OrderSrv interface {
	Get(ctx context.Context, orderSn string) (string, error)
}

type GoodsSrv interface {
	Post(ctx context.Context, orderSn string) (string, error)
}

type orderService struct {
	data string
}

func (o *orderService) Get(ctx context.Context, orderSn string) (string, error) {
	fmt.Println("11111")
	return o.data, nil
}

func newOrderService(sv *service) *orderService {
	return &orderService{
		data: sv.data,
	}
}

type goodsService struct {
	data string
}

func (g *goodsService) Post(ctx context.Context, orderSn string) (string, error) {
	fmt.Println("22222")
	return g.data, nil
}

func newGoodsService(sv *service) *goodsService {
	return &goodsService{
		data: sv.data,
	}
}

var _ OrderSrv = &orderService{}
var _ GoodsSrv = &goodsService{}

// ServiceFactory ----------------------------
type ServiceFactory interface {
	Orders() OrderSrv
	Goods() GoodsSrv
}

type service struct {
	data string
}

func (s *service) Orders() OrderSrv {
	return newOrderService(s)
}

func (s *service) Goods() GoodsSrv {
	return newGoodsService(s)
}

func NewService(data string) *service {
	return &service{data: data}
}

type server struct {
	srv ServiceFactory
}

func NewOrderServer(srv ServiceFactory) *server {
	return &server{srv: srv}
}

func main() {
	srv := NewService("test")
	orderSrv := NewOrderServer(srv)
	orderSrv.srv.Orders().Get(context.Background(), "test")
}

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/config"
	"2106A-zg6/orderAndGoods/goodsproject/order/orderclient"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"
)

type ServiceContext struct {
	Config    config.Config
	GoodsRpc  goodsclient.Goods
	OrdersRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		GoodsRpc:  goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
		OrdersRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrdersRpc)),
	}
}

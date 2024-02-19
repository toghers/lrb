package logic

import (
	"context"

	"2106A-zg6/orderAndGoods/goodsproject/order/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrdersLogic {
	return &UpdateOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrdersLogic) UpdateOrders(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	// todo: add your logic here and delete this line

	return &order.UpdateResponse{}, nil
}

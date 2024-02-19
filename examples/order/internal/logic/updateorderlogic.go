package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/order/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/order/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.UpOrderStatus) (*order.UpOrderStatusResp, error) {

	isCheck := model.UpdateOrderStatus(in.Data.OrderNo, int8(in.Data.Status))

	if !isCheck {
		return nil, status.Error(codes.Internal, "订单状态修改失败")
	}

	return &order.UpOrderStatusResp{}, nil
}

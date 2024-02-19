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

type CreateOrderRollLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRollLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRollLogic {
	return &CreateOrderRollLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRollLogic) CreateOrderRoll(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {

	err := model.DelOrderRoll(in.Data.OrderNo)

	if err != nil {
		return nil, status.Error(codes.Internal, "订单删除")
	}

	return &order.CreateOrderResp{}, nil
}

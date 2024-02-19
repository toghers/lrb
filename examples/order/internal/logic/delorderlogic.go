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

type DelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOrderLogic {
	return &DelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOrderLogic) DelOrder(in *order.DeleteRequest) (*order.DeleteResponse, error) {

	err := model.DelOrder(int(in.Data.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, "删除失败")
	}

	return &order.DeleteResponse{}, nil
}

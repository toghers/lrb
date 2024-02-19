package logic

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/order/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/order/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {

	floatPrice, _ := strconv.ParseFloat(in.Data.Price, 64)

	err, m := model.InsertOrder(in.Data.OrderNo, int(in.Data.UserId), int8(in.Data.Status), floatPrice)
	if err != nil {
		return nil, status.Error(codes.Aborted, "创建失败")
	}

	return &order.CreateOrderResp{
		Msg: &order.OrderInfo{
			UserId:  int64(m.UserId),
			OrderNo: in.Data.OrderNo,
			Price:   in.Data.Price,
			Status:  0,
		},
	}, nil
}

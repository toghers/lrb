package logic

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/order/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/order/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectOrderLogic {
	return &SelectOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectOrderLogic) SelectOrder(in *order.SelectRequest) (*order.SelectResponse, error) {
	page := in.Page
	limit := in.Limit

	err, g := model.LimitAllGood(int(page), int(limit))
	if err != nil {
		return nil, status.Error(codes.Internal, "搜索失败")
	}

	m := []*order.OrderInfo{}
	for _, v := range g {
		strprice := fmt.Sprintf("%.2f", v.TotalPrice)
		m = append(m, &order.OrderInfo{
			Id:      int64(v.Id),
			UserId:  int64(v.UserId),
			OrderNo: v.OrderNo,
			Price:   strprice,
			Status:  int32(v.IsStatus),
		})
	}

	return &order.SelectResponse{
		Data: m,
	}, nil
}

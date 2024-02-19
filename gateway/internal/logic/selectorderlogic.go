package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"
)

type SelectOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectOrderLogic {
	return &SelectOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectOrderLogic) SelectOrder(req *types.SelectOrder) (resp *types.SelectOrderResp, err error) {

	page, _ := strconv.Atoi(req.Page)
	limit, _ := strconv.Atoi(req.Limit)
	res, err := l.svcCtx.OrdersRpc.SelectOrder(l.ctx, &order.SelectRequest{
		Page:  int64(page),
		Limit: int64(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}

	m := []types.Orderinfo{}
	for _, v := range res.Data {
		floatPrice, _ := strconv.ParseFloat(v.Price, 64)
		m = append(m, types.Orderinfo{
			Id:         int(v.Id),
			UserId:     int(v.UserId),
			IsStatus:   int8(v.Status),
			TotalPrice: floatPrice,
			OrderNo:    v.OrderNo,
		})
	}

	return &types.SelectOrderResp{
		Data: m,
	}, nil
}

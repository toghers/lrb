package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"
)

type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrder) (resp *types.DeleteOrderResp, err error) {

	_, err = l.svcCtx.OrdersRpc.DelOrder(l.ctx, &order.DeleteRequest{
		Data: &order.OrderInfo{Id: int64(req.Data.Id)},
	})
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}

	return &types.DeleteOrderResp{
		Code: 200,
		Msg:  "删除成功",
	}, nil
}

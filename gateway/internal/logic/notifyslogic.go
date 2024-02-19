package logic

import (
	"context"
	"fmt"
	"net/url"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/pay"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNotifysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifysLogic {
	return &NotifysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NotifysLogic) Notifys(req *types.NotifyRequest, formData url.Values) (resp *types.NotifyResponse, err error) {

	outTraNo, status, err := pay.NotifyUpdateStatus(formData)

	if err != nil {
		return nil, fmt.Errorf("失败")
	}

	l.svcCtx.OrdersRpc.UpdateOrder(l.ctx, &order.UpOrderStatus{
		Data: &order.OrderInfo{
			OrderNo: outTraNo,
			Status:  int32(status),
		},
	})

	return
}

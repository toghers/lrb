package logic

import (
	"context"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"

	dtmgrpc "github.com/dtm-labs/client/dtmgrpc"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/pay"
	"2106A-zg6/orderAndGoods/goodsproject/order/pb/order"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"
)

type CreateOrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrLogic {
	return &CreateOrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func check(req *types.CreateOrder) error {

	if len(req.Goods) <= 0 {
		return fmt.Errorf("不可没有商品信息")
	}
	return nil

}

func (l *CreateOrLogic) CreateOr(req *types.CreateOrder) (resp *types.CreateOrderResp, err error) {

	//检查参数
	err = check(req)
	if err != nil {
		return nil, err
	}

	arr := []string{}
	goodMap := map[string]string{}

	for _, val := range req.Goods {
		arr = append(arr, val.GoodsId)
		goodMap[val.GoodsId] = val.Num
	}

	//获取商品的信息
	res, err := l.svcCtx.GoodsRpc.GetGoodsById(l.ctx, &goodsclient.SelectByIdReq{
		Id: arr,
	})

	//计算该商品信息的价格
	var count decimal.Decimal
	for _, v := range res.Good {
		fromString, _ := decimal.NewFromString(v.GoodsPrice)

		numStr, ok := goodMap[v.GoodsId]
		if !ok {
			return nil, fmt.Errorf("商品不存在")
		}

		newFromString, _ := decimal.NewFromString(numStr)
		mul := fromString.Mul(newFromString)
		count = count.Add(mul)
	}
	fmt.Println("我是价钱：", count)
	//订单号
	s, _ := uuid.NewV4()
	orderNo := s.String()

	s, _ = uuid.NewV4()

	//将对应的商品信息
	var g []*goods.GoodsStock
	for _, v := range req.Goods {
		g = append(g, &goodsclient.GoodsStock{
			Id:  v.GoodsId,
			Num: v.Num,
		})
	}

	//扣除库存
	goodsStockInfo := &goods.DeductStockReq{
		Data: g,
	}

	orderCreateReq := &order.CreateOrderReq{
		Data: &order.OrderInfo{
			UserId:  1,
			OrderNo: orderNo,
			Price:   count.String(),
			Status:  0,
		},
	}

	payUrl := pay.AlipayGoods(l.svcCtx, "购买商品", orderNo, count.String())

	saga := dtmgrpc.NewSagaGrpc("10.2.176.97:36790", s.String()).
		Add("10.2.176.97:8081/goods.Goods/DeductStock", "10.2.176.97:8081/goods.Goods/DeductStockRollback", goodsStockInfo).
		Add("10.2.176.97:8080/order.Order/CreateOrder", "10.2.176.97:8080/order.Order/CreateOrderRollback", orderCreateReq)

	saga.WaitResult = true
	err = saga.Submit()

	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResp{
		PayUrl: payUrl,
	}, nil
}

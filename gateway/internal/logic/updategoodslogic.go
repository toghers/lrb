package logic

import (
	"context"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(req *types.UpdateGoodRequest) (resp *types.UpdateGoodResponse, err error) {

	m := goodsclient.GoodsResp{
		GoodsId:    req.GoodsOneInfo.GoodsId,
		GoodsName:  req.GoodsOneInfo.GoodsName,
		GoodsPrice: req.GoodsOneInfo.GoodsPrice,
		GoodsNum:   req.GoodsOneInfo.GoodsNum,
		GoodsImage: req.GoodsOneInfo.GoodsImage,
	}

	res, _ := l.svcCtx.GoodsRpc.UpdGoods(l.ctx, &goodsclient.UpdateGoodsInfo{
		Data: &m,
	})

	g := BaseGood(res.GoodsName, res.GoodsNum, res.GoodsImage, res.GoodsRef, res.GoodsPrice, res.GoodsId)

	return &types.UpdateGoodResponse{
		GoodsOneInfo: g,
	}, nil
	return
}

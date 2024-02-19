package logic

import (
	"context"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertGoodsLogic {
	return &InsertGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertGoodsLogic) InsertGoods(req *types.CreateGoodRequest) (resp *types.CreateGoodResponse, err error) {

	m := goodsclient.GoodsResp{
		GoodsName:  req.GoodsOneInfo.GoodsName,
		GoodsPrice: req.GoodsOneInfo.GoodsPrice,
		GoodsNum:   req.GoodsOneInfo.GoodsNum,
		GoodsImage: req.GoodsOneInfo.GoodsImage,
	}

	res, _ := l.svcCtx.GoodsRpc.InsGoods(l.ctx, &goodsclient.CreateGoods{
		Data: &m,
	})

	g := BaseGood(res.Data.GoodsName, res.Data.GoodsNum, res.Data.GoodsImage, res.Data.GoodsRef, res.Data.GoodsPrice, res.Data.GoodsId)

	return &types.CreateGoodResponse{
		GoodsOneInfo: g,
	}, nil

}

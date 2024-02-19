package logic

import (
	"context"
	"strconv"

	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdGoodsLogic {
	return &UpdGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdGoodsLogic) UpdGoods(in *goods.UpdateGoodsInfo) (*goods.GoodsResp, error) {
	// todo: add your logic here and delete this line

	id := in.Data.GoodsId
	goodsName := in.Data.GoodsName
	goodsPrice := in.Data.GoodsPrice
	floatPrice, _ := strconv.ParseFloat(goodsPrice, 64)

	err, g := model.UpdateGood(id, goodsName, floatPrice)
	if err != nil {
		return nil, status.Error(301, "修改失败")
	}

	return &goods.GoodsResp{
		GoodsName:  g.GoodsName,
		GoodsPrice: goodsPrice,
		GoodsNum:   goodsName,
		GoodsRef:   g.GoodsRef,
		GoodsId:    id,
		GoodsImage: g.GoodsImage,
	}, nil
}

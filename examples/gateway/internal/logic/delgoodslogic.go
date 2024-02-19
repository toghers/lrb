package logic

import (
	"context"
	"strconv"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGoodsLogic {
	return &DelGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelGoodsLogic) DelGoods(req *types.DelGoodRequest) (resp *types.DelGoodResponse, err error) {

	delGoodId, _ := strconv.Atoi(req.Goodsid)

	_, err = l.svcCtx.GoodsRpc.DelGoods(l.ctx, &goodsclient.DelGoodsInfo{
		Id: int64(delGoodId),
	})

	if err != nil {
		return &types.DelGoodResponse{
			Code: 301,
			Msg:  "删除失败",
		}, nil
	}

	return &types.DelGoodResponse{
		Code: 200,
		Msg:  "删除成功",
	}, nil

}

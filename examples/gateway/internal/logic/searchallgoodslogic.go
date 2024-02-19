package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/share"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"
)

type SearchAllGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchAllGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllGoodsLogic {
	return &SearchAllGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchAllGoodsLogic) SearchAllGoods(req *types.SearchRequest) (resp *types.SearchResponse, err error) {

	res, _ := l.svcCtx.GoodsRpc.SearchGood(l.ctx, &goodsclient.SearchRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})

	err = share.InitEs(l.svcCtx.Config, res.Data) //添加至es中

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	contexts := req.Context
	atoInt, _ := strconv.Atoi(req.Limit)
	pageInt, _ := strconv.Atoi(req.Page)

	err, m := share.HighSearch(l.svcCtx.Config, contexts, pageInt, atoInt) //高亮搜索

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &types.SearchResponse{
		Goodinfo: m,
	}, nil

	//var m []types.Goodinfo
	//for _, val := range res.Data {
	//	m = append(m, types.Goodinfo{
	//		GoodsName:  val.GoodsName,
	//		GoodsPrice: val.GoodsPrice,
	//		GoodsNum:   val.GoodsNum,
	//		GoodsRef:   val.GoodsRef,
	//		GoodsId:    val.GoodsId,
	//		GoodsImage: val.GoodsImage,
	//	})
	//}

}

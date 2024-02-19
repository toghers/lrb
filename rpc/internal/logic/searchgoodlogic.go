package logic

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodLogic {
	return &SearchGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGoodLogic) SearchGood(in *goods.SearchRequest) (*goods.SearchResponse, error) {

	page := in.Page
	limit := in.Limit
	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)

	err, g := model.LimitAllGood(intPage, intLimit)
	if err != nil {
		return nil, status.Error(codes.Internal, "搜索失败")
	}

	var m []*goods.GoodsResp
	for _, val := range g {
		strPrice := fmt.Sprintf("%.2f", val.GoodsPrice)
		m = append(m, &goods.GoodsResp{
			GoodsName:  val.GoodsName,
			GoodsPrice: strPrice,
			GoodsNum:   strconv.Itoa(val.GoodsNum),
			GoodsRef:   val.GoodsRef,
			GoodsId:    strconv.Itoa(val.Id),
			GoodsImage: val.GoodsImage,
		})
	}

	return &goods.SearchResponse{
		Data: m,
	}, nil
}

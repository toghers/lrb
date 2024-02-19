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

type SelGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelGoodsLogic {
	return &SelGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func SelectCheck(in *goods.SelectGoodsInfo) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "参数不可为空")
	}
	if in.Id <= 0 {
		return status.Error(codes.InvalidArgument, "id is not null")
	}
	return nil

}

func (l *SelGoodsLogic) SelGoods(in *goods.SelectGoodsInfo) (*goods.SelectGoodsResp, error) {
	err := SelectCheck(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "参数不可不传")
	}

	id := in.Id

	err, g := model.SelGood(id)
	if err != nil {
		return nil, status.Error(codes.Internal, "失败")
	}

	strPrice := fmt.Sprintf("%.2f", g.GoodsPrice)

	info, _ := BaseInfo(g.GoodsName, strPrice, strconv.Itoa(g.GoodsNum), g.GoodsRef, g.GoodsImage, strconv.Itoa(g.Id))

	return &goods.SelectGoodsResp{
		Data: info,
	}, nil
}

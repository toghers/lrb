package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsGoodsLogic {
	return &InsGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checks(in *goods.CreateGoods) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "参数不可为空")
	}

	if in.Data.GoodsName == "" {
		return status.Error(codes.InvalidArgument, "goodsName is required")
	}

	if in.Data.GoodsNum == "" {
		return status.Error(codes.InvalidArgument, "GoodsNum is required")
	}

	if in.Data.GoodsImage == "" {
		return status.Error(codes.InvalidArgument, "GoodsImage is required")
	}

	if in.Data.GoodsPrice == "" {
		return status.Error(codes.InvalidArgument, "GoodsPrice is required")
	}

	price, err := decimal.NewFromString(in.Data.GoodsPrice)
	if err != nil {
		return err
	}

	if price.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("good price must > 0")
	}
	return nil
}

func (l *InsGoodsLogic) InsGoods(in *goods.CreateGoods) (*goods.CreateGoodsResp, error) {

	err := checks(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "参数不可不传")
	}

	name := in.Data.GoodsName
	num := in.Data.GoodsNum
	price := in.Data.GoodsPrice
	image := in.Data.GoodsImage

	isCheck, g := model.InsertGood(name, price, num, image)
	if isCheck != true {
		return nil, status.Error(codes.Internal, "添加失败")
	}

	info, _ := BaseInfo(name, price, num, g.GoodsRef, image, strconv.Itoa(g.Id))
	return &goods.CreateGoodsResp{Data: info}, nil

}

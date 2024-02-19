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

type GetGoodsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdLogic {
	return &GetGoodsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIdLogic) GetGoodsById(in *goods.SelectByIdReq) (*goods.SelectByIdResp, error) {
	// todo: add your logic here and delete this line

	ids := in.Id

	arr := []int{}
	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		arr = append(arr, id)
	}

	Ischeck, g := model.SelectInfoById(arr)
	if !Ischeck {
		return nil, status.Error(codes.Internal, "查找")
	}

	respGoods := []*goods.GoodsResp{}
	for _, v := range g {
		strprice := fmt.Sprintf("%.2f", v.GoodsPrice)
		strNum := fmt.Sprintf("%d", v.GoodsNum)
		strid := fmt.Sprintf("%d", v.Id)
		respGoods = append(respGoods, &goods.GoodsResp{
			GoodsName:  v.GoodsName,
			GoodsPrice: strprice,
			GoodsNum:   strNum,
			GoodsRef:   v.GoodsRef,
			GoodsId:    strid,
			GoodsImage: v.GoodsImage,
		})
	}

	return &goods.SelectByIdResp{
		Good: respGoods,
	}, nil
}

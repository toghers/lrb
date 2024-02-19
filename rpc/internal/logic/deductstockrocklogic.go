package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockRockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockRockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockRockLogic {
	return &DeductStockRockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockRockLogic) DeductStockRock(in *goods.DeductStockReq) (*goods.DeductStockResp, error) {

	isCheck := model.UpdateStockRoll(in.Data)
	if isCheck != nil {
		return nil, status.Error(codes.Internal, "修改失败")
	}

	return &goods.DeductStockResp{}, nil
}

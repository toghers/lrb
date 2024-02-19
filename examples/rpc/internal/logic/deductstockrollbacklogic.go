package logic

import (
	"context"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductStockRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductStockRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductStockRollbackLogic {
	return &DeductStockRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductStockRollbackLogic) DeductStockRollback(in *goods.DeductStockReq) (*goods.DeductStockResp, error) {
	// todo: add your logic here and delete this line

	return &goods.DeductStockResp{}, nil
}

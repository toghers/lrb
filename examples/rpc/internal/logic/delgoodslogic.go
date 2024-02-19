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

type DelGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGoodsLogic {
	return &DelGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func check(in *goods.DelGoodsInfo) error {
	if in == nil {
		return status.Error(codes.InvalidArgument, "参数不可为空")
	}
	if in.Id <= 0 {
		return status.Error(codes.InvalidArgument, "id is not null")
	}
	return nil

}
func (l *DelGoodsLogic) DelGoods(in *goods.DelGoodsInfo) (*goods.EmptyResp, error) {

	err := check(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "参数错误")
	}

	id := in.Id
	isCheck := model.DelGood(int(id))

	if isCheck != nil {
		return nil, status.Error(codes.Internal, "删除失败")
	}

	return &goods.EmptyResp{}, nil
}

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

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *goods.LoginRequest) (*goods.LoginResponse, error) {
	username := in.Username

	isCheck, u := model.UserLogin(username)

	if !isCheck {
		return nil, status.Error(codes.Internal, "没有该用户")
	}

	return &goods.LoginResponse{
		Id:       int64(u.Id),
		Password: u.Password,
	}, nil
}

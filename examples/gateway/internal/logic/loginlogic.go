package logic

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/share"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/goodsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Username == "" {
		return &types.LoginResponse{
			Code: 301,
			Msg:  "用户名不为空",
		}, nil
	}

	if req.Password == "" {
		return &types.LoginResponse{
			Code: 301,
			Msg:  "密码不可为空",
		}, nil
	}

	res, err := l.svcCtx.GoodsRpc.Login(l.ctx, &goodsclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResponse{
			Code: 302,
			Msg:  "登录失败",
		}, nil
	}

	if req.Password != res.Password {
		return &types.LoginResponse{
			Code: 303,
			Msg:  "账号或密码错误",
		}, nil
	}

	tokens := share.CreateTokens(&share.MyCustomClaims{
		Id: int(res.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire)),
		},
	}, []byte(l.svcCtx.Config.Auth.AccessSecret))

	return &types.LoginResponse{
		Code:  200,
		Msg:   "登陆成功",
		Token: tokens,
	}, nil
}

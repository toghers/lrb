package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
	GoodsRpc  zrpc.RpcClientConf
	OrdersRpc zrpc.RpcClientConf
	Alipays   struct {
		AppId        string
		PrivateKey   string
		AliPublicKey string
		NotifyUrl    string
	}
	Esurl string
}

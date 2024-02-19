package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	NacInfo struct {
		Host  string
		Port  int
		Space string
	}
}

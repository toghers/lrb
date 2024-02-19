package main

import (
	"flag"
	"fmt"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/initialization"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/config"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/model"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/server"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	initialization.InitNa(c) //初始化nacos
	model.InitMysql()        //初始化mysql
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, server.NewGoodsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

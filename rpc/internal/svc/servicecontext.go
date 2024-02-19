package svc

import (
	"2106A-zg6/orderAndGoods/goodsproject/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

//func InitModels(c config.Config) sqlx.SqlConn {
//	cfg := c.Mysql
//	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfg.User, cfg.Pwd, cfg.Host, cfg.Port, cfg.Database)
//	db := sqlx.NewMysql(dsn)
//	_, err := db.Exec("select (1)")
//	if err != nil {
//		panic(err)
//	}
//	return db
//
//}

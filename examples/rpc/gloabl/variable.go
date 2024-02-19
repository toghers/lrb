package gloabl

import "gorm.io/gorm"

var (
	Db        *gorm.DB
	Conf      Configs
	NocosConf Yamls
)

type Configs struct {
	Mysql struct {
		Dsn string
	}
}

type Yamls struct {
	NacInfo struct {
		Host  string
		Port  int
		Space string
	}
}

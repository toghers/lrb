package gloabl

import "gorm.io/gorm"

var (
	Db     *gorm.DB
	Config Configs
)

type Configs struct {
	Mysql struct {
		Dsn string
	}
}

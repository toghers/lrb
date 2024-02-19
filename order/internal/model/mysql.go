package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"2106A-zg6/orderAndGoods/goodsproject/order/gloabl"
)

type Goods struct {
	Id         int
	GoodsName  string
	GoodsPrice float64
	GoodsNum   int
	GoodsRef   string
	IsDel      int8
	Status     int8
	CreateAt   int
	UpdateAt   int
	DeleteAt   int
	GoodsImage string
	IsHot      int8
	Stock      int
}

type Order struct {
	Id         int
	UserId     int
	GoodsId    int
	IsStatus   int8
	IsDel      int8
	TotalPrice float64
	DeleteAt   int
	CreateAt   int
	UpdateAt   int
	OrderNo    string
}

type Users struct {
	Id       int
	Username string
	Password string
	Mobile   string
	CreateAt int
	UpdateAt int
	DeleteAt int
	IsDel    int8
	Score    int
	Actives  string
}

func InitMysql() {
	dsn := gloabl.Config.Mysql.Dsn
	var err error
	gloabl.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		return
	}
}

func InsertOrder(orderNo string, userid int, status int8, totalPrice float64) (error, Order) {

	o := Order{
		UserId:     userid,
		GoodsId:    0,
		IsStatus:   status,
		IsDel:      0,
		TotalPrice: totalPrice,
		DeleteAt:   0,
		CreateAt:   0,
		UpdateAt:   0,
		OrderNo:    orderNo,
	}
	res := gloabl.Db.Table("order").Create(&o)
	if res.Error != nil {
		return res.Error, Order{}
	}
	return nil, o

}

func DelOrderRoll(orderNo string) error {

	var g Order

	res := gloabl.Db.Table("order").Where("order_no=?", orderNo).First(&g)
	if res.Error != nil {
		return res.Error
	}

	g.IsDel = 1

	res = gloabl.Db.Table("order").Save(&g)
	if res.Error != nil {
		return res.Error
	}

	return res.Error

}

func UpdateOrderStatus(orderTo string, status int8) bool {

	UpDateUnix := time.Now().Unix() //修改时间
	log.Println("我是订单：", orderTo)
	res := gloabl.Db.Table("order").Where("order_no = ?", orderTo).Updates(Order{IsStatus: status, UpdateAt: int(UpDateUnix)})

	//res := gloabl.Db.Model(&Order{}).Select("order_no", "update_at").Updates(Order{IsStatus: status, UpdateAt: int(UpDateUnix)})
	//res := gloabl.Db.Table("order").Where("order_no = ?", orderTo).Update("is_status", status)

	if res.Error != nil {
		return false
	}

	return true
}

func DelOrder(id int) error {
	var g Order

	res := gloabl.Db.Table("order").Where("id=?", id).First(&g)
	if res.Error != nil {
		return res.Error
	}

	g.IsDel = 1

	res = gloabl.Db.Table("order").Save(&g)
	if res.Error != nil {
		return res.Error
	}

	return res.Error
}

func LimitAllGood(page int, limit int) (error, []Order) {

	offset := (page - 1) * limit

	g := []Order{}

	res := gloabl.Db.Table("order").Where("is_del!=?", 1).Offset(offset).Limit(limit).Find(&g)

	if res.Error != nil {
		return res.Error, []Order{}
	}

	return nil, g
}

package model

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"2106A-zg6/orderAndGoods/goodsproject/rpc/gloabl"
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"
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
	Id       int
	UserId   int
	GoodsId  int
	IsStatus int8
	CreateAt int
	DeleteAt int
	IsDel    int8
}

type Stocks struct {
	Id       int
	Stock    int
	GoodsId  int
	CreateAt int
	UpdateAt int
	DeleteAt int
	IsDel    int8
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

	fmt.Println(gloabl.Conf.Mysql.Dsn)
	dsn := gloabl.Conf.Mysql.Dsn
	var err error
	gloabl.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		return
	}
}

func InsertGood(name string, price string, num string, image string) (bool, Goods) {

	floatPrices, _ := strconv.ParseFloat(price, 64)
	intNum, _ := strconv.Atoi(num)

	//商品编号
	TimesUnix := time.Now().UnixMicro()
	strTimeUnix := strconv.FormatInt(TimesUnix, 10)
	GoodsRefInfo := "sps" + strTimeUnix
	//创建时间
	CreateTimes := time.Now().Unix()
	m := Goods{
		GoodsName:  name,
		GoodsPrice: floatPrices,
		GoodsNum:   intNum,
		GoodsRef:   GoodsRefInfo,
		IsDel:      0,
		Status:     0,
		CreateAt:   int(CreateTimes),
		UpdateAt:   int(CreateTimes),
		DeleteAt:   0,
		GoodsImage: image,
		IsHot:      0,
	}
	res := gloabl.Db.Table("goods").Create(&m)
	if res.Error != nil {
		return false, Goods{}
	}
	return true, m

}

func DelGood(id int) error {
	var g Goods

	res := gloabl.Db.Table("goods").Where("id=?", id).First(&g)
	if res.Error != nil {
		return res.Error
	}

	g.IsDel = 1

	res = gloabl.Db.Table("goods").Save(&g)
	if res.Error != nil {
		return res.Error
	}

	return res.Error
}

func UpdateGood(id string, goodsName string, floatPrice float64) (error, Goods) {
	var g Goods
	res := gloabl.Db.Table("goods").Where("id=?", id).First(&g)

	if res.Error != nil {
		return res.Error, Goods{}
	}

	g.GoodsName = goodsName
	g.GoodsPrice = floatPrice

	res = gloabl.Db.Table("goods").Where("id=?", id).Save(&g)
	if res.Error != nil {
		return res.Error, Goods{}
	}

	return nil, g

}

func SelGood(id int64) (error, Goods) {
	var g Goods

	res := gloabl.Db.Table("goods").Where("id=?", id).Where("is_del != ?", 1).First(&g)
	if res.Error != nil {
		return res.Error, Goods{}
	}

	return nil, g

}

func LimitAllGood(page int, limit int) (error, []Goods) {
	offset := (page - 1) * limit

	g := []Goods{}

	res := gloabl.Db.Table("goods").Where("is_del!=?", 1).Offset(offset).Limit(limit).Find(&g)

	if res.Error != nil {
		return res.Error, []Goods{}
	}

	return nil, g
}

func UserLogin(username string) (bool, Users) {
	var u Users

	res := gloabl.Db.Table("users").Where("username=?", username).First(&u)

	if res.Error != nil {
		return false, Users{}
	}

	return true, u

}

func SelectInfoById(id []int) (bool, []Goods) {

	g := []Goods{}
	res := gloabl.Db.Table("goods").Where("id in (?)", id).Find(&g)

	if res.Error != nil {
		return false, []Goods{}
	}

	return true, g

}

func UpdateStock(data []*goods.GoodsStock) error {

	tx := gloabl.Db.Begin()

	for _, v := range data {

		var good Goods

		err := tx.Where("id=?", v.Id).First(&good).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		intNum, _ := strconv.Atoi(v.Num)
		if good.Stock < intNum {
			tx.Rollback()
			return errors.New("库存不足")
		}

		good.Stock -= intNum
		err = tx.Model(&Goods{}).Where("id=?", v.Id).Update("stock", good.Stock).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func UpdateStockRoll(data []*goods.GoodsStock) error {

	tx := gloabl.Db.Begin()
	for _, v := range data {
		var good Goods
		err := tx.Where("id=?", v.Id).First(&good).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		intNum, _ := strconv.Atoi(v.Num)

		good.Stock += intNum
		err = tx.Model(&Goods{}).Where("id=?", v.Id).Update("stock", good.Stock).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil

}

package logic

import (
	"2106A-zg6/orderAndGoods/goodsproject/rpc/pb/goods"
)

func BaseInfo(name, price, num, ref, image string, id string) (*goods.GoodsResp, error) {
	return &goods.GoodsResp{
		GoodsName:  name,
		GoodsPrice: price,
		GoodsNum:   num,
		GoodsRef:   ref,
		GoodsId:    id,
		GoodsImage: image,
	}, nil
}

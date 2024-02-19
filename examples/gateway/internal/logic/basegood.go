package logic

import "2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"

func BaseGood(name, num, image, ref string, price string, id string) types.Goodinfo {

	m := types.Goodinfo{
		GoodsName:  name,
		GoodsPrice: price,
		GoodsNum:   num,
		GoodsRef:   ref,
		GoodsId:    id,
		GoodsImage: image,
	}
	return m

}

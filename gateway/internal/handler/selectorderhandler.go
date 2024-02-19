package handler

import (
	"net/http"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/logic"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SelectOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectOrder
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSelectOrderLogic(r.Context(), svcCtx)
		resp, err := l.SelectOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

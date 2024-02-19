package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/logic"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/svc"
	"2106A-zg6/orderAndGoods/goodsproject/gateway/internal/types"
)

func NotifysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewNotifysLogic(r.Context(), svcCtx)

		r.ParseForm()
		_, err := l.Notifys(&req, r.PostForm)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			_, err = w.Write([]byte("success"))
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			httpx.Ok(w)
		}
	}
}

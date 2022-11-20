package handler

import (
	"net/http"

	"cloudclass_go/internal/logic"
	"cloudclass_go/internal/svc"
	"cloudclass_go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func kebiaochangeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Kebiaocreq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewKebiaochangeLogic(r.Context(), svcCtx)
		resp, err := l.Kebiaochange(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

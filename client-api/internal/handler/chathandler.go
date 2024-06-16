package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"langriver-client/internal/logic"
	"langriver-client/internal/svc"
	"langriver-client/internal/types"
)

func chatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatLogic(r.Context(), svcCtx)
		resp, err := l.Chat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

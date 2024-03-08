package user

import (
	"go-gin/internal/errorx"
	"go-gin/logic/user"
	"go-gin/types"
	"go-gin/utils/httpx"

	"github.com/gin-gonic/gin"
)

func ListUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListUserReq
		if err := ctx.ShouldBind(&req); err != nil {
			httpx.Error(ctx, errorx.NewWithError(err))
			return
		}

		l := user.NewListUser()
		resp, err := l.Handle(req)
		if err != nil {
			httpx.Error(ctx, err)
			return
		}
		httpx.Ok(ctx, resp)

	}
}

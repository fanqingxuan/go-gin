package user

import (
	"go-gin/internal/errorx"
	"go-gin/types"
	"go-gin/utils/httpx"

	"github.com/gin-gonic/gin"
)

func AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddUserReq
		if err := ctx.ShouldBind(&req); err != nil {
			httpx.Error(ctx, errorx.NewWithError(err))
			return
		}
		httpx.Ok(ctx, types.AddUserReply{
			Message: "hello" + req.Name,
		})

	}
}

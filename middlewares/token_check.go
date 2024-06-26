package middlewares

import (
	"go-gin/consts"
	"go-gin/internal/ginx/httpx"

	"github.com/gin-gonic/gin"
)

type TokenHeader struct {
	Token string `header:"token" binding:"required"`
}

func TokenCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req TokenHeader
		if err := ctx.ShouldBindHeader(&req); err != nil {
			httpx.Error(ctx, consts.ErrUserMustLogin)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

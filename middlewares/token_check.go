package middlewares

import (
	"go-gin/consts"
	"go-gin/internal/errorx"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/token"

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
		if has, err := token.Has(ctx, req.Token); err != nil {
			httpx.Error(ctx, errorx.NewDefault("获取token错误"))
			ctx.Abort()
			return
		} else if !has {
			httpx.Error(ctx, consts.ErrUserNeedLoginAgain)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

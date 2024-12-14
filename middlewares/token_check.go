package middlewares

import (
	"go-gin/consts"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
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
			httpx.Error(httpx.NewContext(ctx), consts.ErrUserMustLogin)
			ctx.Abort()
		}
		if has, err := token.Has(ctx, req.Token); err != nil {
			httpx.Error(httpx.NewContext(ctx), errorx.NewDefault("获取token错误"))
			ctx.Abort()
		} else if !has {
			httpx.Error(httpx.NewContext(ctx), consts.ErrUserNeedLoginAgain)
			ctx.Abort()
		}
		ctx.Next()
	}
}

package middlewares

import (
	"go-gin/consts"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
	"go-gin/internal/token"
)

type TokenHeader struct {
	Token string `header:"token" binding:"required"`
}

func TokenCheck() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (interface{}, error) {
		var req TokenHeader
		if err := ctx.ShouldBindHeader(&req); err != nil {
			httpx.Error(ctx, consts.ErrUserMustLogin)
			ctx.Abort()
			return nil, err
		}
		if has, err := token.Has(ctx, req.Token); err != nil {
			return nil, errorx.NewDefault("获取token错误")
		} else if !has {
			return nil, consts.ErrUserNeedLoginAgain
		}
		return nil, nil
	}
}

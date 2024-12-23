package middleware

import (
	"go-gin/const/errcode"
	"go-gin/internal/httpx"
	"go-gin/internal/token"
)

type TokenHeader struct {
	Token string `header:"token" binding:"required"`
}

func TokenCheck() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (any, error) {
		var req TokenHeader
		if err := ctx.ShouldBindHeader(&req); err != nil {
			return nil, errcode.ErrUserMustLogin
		}
		if has, err := token.Has(ctx, req.Token); err != nil {
			return nil, errcode.NewDefault("获取token错误")
		} else if !has {
			return nil, errcode.ErrUserNeedLoginAgain
		}
		return nil, nil
	}
}

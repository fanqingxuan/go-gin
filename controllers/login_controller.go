package controllers

import (
	"go-gin/internal/components/logx"
	"go-gin/internal/httpx"
	"go-gin/internal/token"
	"go-gin/logic"
	"go-gin/types"
)

type loginController struct {
}

var LoginController = &loginController{}

func (c *loginController) Login(ctx *httpx.Context) (interface{}, error) {
	l := logic.NewLoginLogic()
	var req types.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		return nil, err
	}
	return l.Handle(ctx, req)
}

func (c *loginController) LoginOut(ctx *httpx.Context) (interface{}, error) {
	token.Flush(ctx, ctx.GetHeader("token"))
	return nil, nil
}

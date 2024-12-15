package controller

import (
	"go-gin/internal/httpx"
	"go-gin/internal/token"
	"go-gin/logic"
)

type loginController struct {
}

var LoginController = &loginController{}

func (c *loginController) Login(ctx *httpx.Context) (interface{}, error) {
	return ShouldBindHandle(ctx, logic.NewLoginLogic())
}

func (c *loginController) LoginOut(ctx *httpx.Context) (interface{}, error) {
	token.Flush(ctx, ctx.GetHeader("token"))
	return nil, nil
}

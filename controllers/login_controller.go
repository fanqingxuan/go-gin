package controllers

import (
	"go-gin/internal/components/logx"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/token"
	"go-gin/logic"
	"go-gin/types"

	"github.com/gin-gonic/gin"
)

type loginController struct {
}

var LoginController = &loginController{}

func (c *loginController) Login(ctx *gin.Context) {
	l := logic.NewLoginLogic()
	var req types.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		httpx.Error(ctx, err)
		return
	}
	resp, err := l.Handle(ctx, req)
	httpx.Handle(ctx, resp, err)
}

func (c *loginController) LoginOut(ctx *gin.Context) {
	token.Flush(ctx, ctx.GetHeader("token"))
	httpx.OkResponse(ctx)
}

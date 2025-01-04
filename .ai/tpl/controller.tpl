package controller

import (
	"go-gin/internal/httpx"
	"go-gin/logic"
)

type xxxController struct {
}

var XxxController = &xxxController{}

func (c *xxxController) TT(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewTTLogic())
}
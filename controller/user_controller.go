package controller

import (
	"go-gin/internal/httpx"
	"go-gin/logic"
)

type userController struct {
}

var UserController = &userController{}

func (c *userController) Index(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewIndexLogic())
}

func (c *userController) List(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewGetUsersLogic())
}

func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewAddUserLogic())
}

func (c *userController) MultiUserAdd(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewMultiAddUserLogic())
}

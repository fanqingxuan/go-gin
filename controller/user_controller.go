package controller

import (
	"go-gin/internal/httpx"
	"go-gin/logic"
	"go-gin/typing"

	"github.com/golang-module/carbon/v2"
)

type userController struct {
}

var UserController = &userController{}

type User struct {
	Name       string        `json:"name"`
	CreateTime carbon.Carbon `json:"create_time"`
}

func (c *userController) Index(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewIndexLogic())
}

func (c *userController) List(ctx *httpx.Context) (any, error) {
	var req typing.ListReq
	l := logic.NewGetUsersLogic()
	return l.Handle(ctx, req)
}

func (c *userController) AddUser(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewAddUserLogic())
}

func (c *userController) MultiUserAdd(ctx *httpx.Context) (any, error) {
	return httpx.ShouldBindHandle(ctx, logic.NewMultiAddUserLogic())
}

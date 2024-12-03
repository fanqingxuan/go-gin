package controllers

import (
	"go-gin/events"
	"go-gin/internal/components/logx"
	"go-gin/internal/event"
	"go-gin/internal/ginx/httpx"
	"go-gin/logic"
	"go-gin/types"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

type userController struct {
}

var UserController = &userController{}

type User struct {
	Name       string        `json:"name"`
	CreateTime carbon.Carbon `json:"create_time"`
}

func (c *userController) Index(ctx *gin.Context) {
	event.Fire(ctx, events.NewSampleEvent("hello 测试"))
	events.NewSampleEvent("333").Fire(ctx)
	u := User{
		Name:       "hello",
		CreateTime: carbon.Parse("now").AddCentury(),
	}
	httpx.Ok(ctx, u)
}

func (c *userController) List(ctx *gin.Context) {
	var req types.ListReq
	l := logic.NewGetUsersLogic()
	resp, err := l.Handle(ctx, req)
	httpx.Handle(ctx, resp, err)
}

func (c *userController) AddUser(ctx *gin.Context) {
	var req types.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		httpx.Error(ctx, err)
		return
	}
	l := logic.NewAddUserLogic()
	resp, err := l.Handle(ctx, req)
	if err != nil {
		httpx.Error(ctx, err)
		return
	}

	httpx.Ok(ctx, resp)
}

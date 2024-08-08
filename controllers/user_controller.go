package controllers

import (
	"go-gin/events"
	"go-gin/internal/components/logx"
	"go-gin/internal/event"
	"go-gin/internal/ginx/httpx"
	"go-gin/services"
	"go-gin/types"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service *services.UserService
}

var UserController = &userController{
	service: services.NewUserService(),
}

func (c *userController) Index(ctx *gin.Context) {
	event.Fire(ctx, events.CreateSampleEvent("hello 测试"))
	// events.CreateSampleEvent("测试").Dispatch(ctx)
	httpx.Ok(ctx, "hello world")
}

func (c *userController) List(ctx *gin.Context) {
	var req types.ListReq
	resp, err := c.service.GetAllUsers(ctx, req)
	httpx.Handle(ctx, resp, err)
}

func (c *userController) AddUser(ctx *gin.Context) {
	var req types.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		httpx.Error(ctx, err)
		return
	}

	resp, err := c.service.AddUser(ctx, req)
	if err != nil {
		httpx.Error(ctx, err)
		return
	}

	httpx.Ok(ctx, resp)
}

package controllers

import (
	"fmt"
	"go-gin/events"
	"go-gin/internal/components/logx"
	"go-gin/internal/event"
	"go-gin/internal/ginx/httpx"
	"go-gin/models"
	"go-gin/services"
	"go-gin/types"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService *services.UserService
}

var UserController = &userController{
	userService: services.NewUserService(),
}

func (c *userController) Index(ctx *gin.Context) {
	event.Fire(ctx, events.CreateSampleEvent("hello 测试"))
	// events.CreateSampleEvent("测试").Dispatch(ctx)
	httpx.Ok(ctx, "hello world")
}

func (c *userController) List(ctx *gin.Context) {
	u, err := c.userService.GetAllUsers(ctx)
	httpx.Handle(ctx, u, err)
}

func (c *userController) AddUser(ctx *gin.Context) {
	var req types.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		logx.WithContext(ctx).Warn("ShouldBind异常", err)
		httpx.Error(ctx, err)
		return
	}
	user := &models.User{
		Name: req.Name,
		Age:  &req.Age,
	}
	err := c.userService.AddUser(ctx, user)
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	resp := types.AddUserReply{
		Message: fmt.Sprintf("add user succcess %s=%d", user.Name, user.Id),
	}
	httpx.Ok(ctx, resp)
}

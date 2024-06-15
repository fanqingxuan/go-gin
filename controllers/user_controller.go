package controllers

import (
	"fmt"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/ginx/validators"
	"go-gin/models"
	"go-gin/pkg/logx"
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
	type User struct {
		Name string `binding:"required,min=5" label:"姓,44名"`
	}
	u := User{Name: "测试们啊"}
	err := validators.Validate(u)
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	logx.WithContext(ctx).Info("门店编程", u)
	logx.WithContext(ctx).Infof("关键字", "这是什么%s/%s", "54", "5444")
	httpx.Ok(ctx, "hello world")
}

func (c *userController) List(ctx *gin.Context) {
	logx.WithContext(ctx).Debug("获悉信息", "这是debug日志")
	logx.WithContext(ctx).Info("获悉信息", "这是debug日志")
	logx.WithContext(ctx).Warn("获悉信息", "这是debug日志")
	logx.WithContext(ctx).Error("获悉信息", "这是debug日志")
	u, err := c.userService.GetAllUsers(ctx)
	httpx.Handle(ctx, u, err)
}

func (c *userController) AddUser(ctx *gin.Context) {
	var req types.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
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

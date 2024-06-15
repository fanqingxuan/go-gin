package controllers

import (
	"fmt"
	"go-gin/models"
	"go-gin/services"
	"go-gin/types"
	"go-gin/utils/errorx"
	"go-gin/utils/httpx"
	"go-gin/utils/validators"

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
		Name string `binding:"required" label:"姓,44名"`
	}
	u := User{}
	err := validators.Validate(u)
	if err != nil {
		httpx.Error(ctx, err)
		return
	}
	httpx.Ok(ctx, "hello world")
}

func (c *userController) List(ctx *gin.Context) {
	u, err := c.userService.GetAllUsers(ctx)
	httpx.Handle(ctx, u, err)
}

func (c *userController) AddUser(ctx *gin.Context) {
	var req types.AddUserReq
	if err := ctx.ShouldBind(&req); err != nil {
		httpx.Error(ctx, errorx.NewDefault(err.Error()))
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

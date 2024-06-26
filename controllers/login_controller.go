package controllers

import (
	"errors"
	"fmt"
	"go-gin/consts"
	"go-gin/internal/components/db"
	"go-gin/internal/errorx"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/token"
	"go-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginController struct {
}

var LoginController = &loginController{}

func (c *loginController) Login(ctx *gin.Context) {
	var user models.User
	if err := db.WithContext(ctx).Find(&user, "username=?", "测试1").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpx.Error(ctx, consts.ErrUserNameOrPwdFaild)
		} else {
			httpx.Error(ctx, err)
		}
		return
	}
	t := token.GetTokenId()

	if err := token.Set(ctx, t, "name", "测试"); err != nil {
		httpx.Error(ctx, errorx.NewDefault("存储用户信息异常"))
		return
	}
	httpx.Ok(ctx, t)
}

func (c *loginController) LoginOut(ctx *gin.Context) {
	fmt.Println()
	token.FlushAll(ctx, ctx.GetHeader("token"))
	httpx.OkResponse(ctx)
}

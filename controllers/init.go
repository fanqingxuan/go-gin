package controllers

import (
	"go-gin/consts"
	"go-gin/utils/httpx"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	route.NoMethod(func(ctx *gin.Context) {
		httpx.Error(ctx, consts.ErrMethodNotAllowed)
	})
	route.NoRoute(func(ctx *gin.Context) {
		httpx.Error(ctx, consts.ErrNoRoute)
	})

	user_router := route.Group("/user")
	user_router.GET("/", UserController.Index)
	user_router.GET("/list", UserController.List)
	user_router.GET("/adduser", UserController.AddUser)

}

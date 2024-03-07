package handler

import (
	"go-gin/internal/errorx"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {

	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(200, "测试")
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, errorx.NewDefault("路由不存在"))
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

}

package handler

import (
	"go-gin/app/common/httpx"
	"go-gin/internal/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {

	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(200, "测试")
	})
	r.NoRoute(func(ctx *gin.Context) {
		httpx.Error(ctx, errorx.New(http.StatusNotFound, "路由不存在"))
	})

	r.GET("/", func(ctx *gin.Context) {
		httpx.Ok(ctx, "世界你好")
	})

}

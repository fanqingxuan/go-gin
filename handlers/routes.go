package handlers

import (
	"go-gin/internal/errorx"
	"go-gin/utils/httpx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {

	r.NoMethod(func(ctx *gin.Context) {
		httpx.Error(ctx, errorx.New(http.StatusMethodNotAllowed, "方法不允许"))
	})
	r.NoRoute(func(ctx *gin.Context) {
		httpx.Error(ctx, errorx.New(http.StatusNotFound, "路由不存在"))
	})

	r.GET("/", func(ctx *gin.Context) {
		httpx.Ok(ctx, "世界你好")
	})

}

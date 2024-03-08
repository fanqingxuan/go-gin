package user

import (
	"go-gin/utils/httpx"

	"github.com/gin-gonic/gin"
)

func AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Ok(ctx, "add页面")

	}
}

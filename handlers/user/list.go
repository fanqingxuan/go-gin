package user

import (
	"go-gin/utils/httpx"

	"github.com/gin-gonic/gin"
)

func ListUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpx.Ok(ctx, "这是list页面")

	}
}

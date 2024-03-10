package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/utils"
)

func traceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := utils.NewUuid()
		ctx.Set("requestId", traceId)
		ctx.Next()
	}
}

package middlewares

import (
	"go-gin/internal/traceid"

	"github.com/gin-gonic/gin"
)

func traceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(traceid.TraceIdFieldName, traceid.New())
		ctx.Next()
	}
}

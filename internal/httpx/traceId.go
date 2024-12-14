package httpx

import (
	"go-gin/internal/traceid"

	"github.com/gin-gonic/gin"
)

func TraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(traceid.TraceIdFieldName, traceid.New())
		ctx.Next()
	}
}

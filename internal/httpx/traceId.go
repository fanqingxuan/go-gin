package httpx

import (
	"context"
	"go-gin/internal/traceid"

	"github.com/gin-gonic/gin"
)

func TraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := traceid.New()
		rctx := context.WithValue(ctx.Request.Context(), traceid.TraceIdFieldName, traceId)
		ctx.Request = ctx.Request.WithContext(rctx)
		ctx.Next()
	}
}

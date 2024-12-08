package httpx

import (
	"go-gin/internal/traceid"
)

func TraceId() HandlerFunc {
	return func(ctx *Context) (interface{}, error) {
		ctx.Set(traceid.TraceIdFieldName, traceid.New())
		ctx.Next()
		return nil, nil
	}
}

package middlewares

import (
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
	"go-gin/internal/utils"
)

func recoverLog() httpx.HandlerFunc {
	return func(ctx *httpx.Context) {
		defer func() {
			if err := recover(); err != nil {
				m := map[string]interface{}{
					"error": fmt.Sprintf("%v", err),
					"file":  utils.FileWithLineNum(),
				}
				logx.WithContext(ctx).Error("panic", m)
				httpx.Error(ctx, errorx.ErrInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()

	}
}

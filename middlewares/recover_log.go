package middlewares

import (
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/errorx"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

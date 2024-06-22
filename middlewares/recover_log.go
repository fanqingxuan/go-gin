package middlewares

import (
	"go-gin/consts"
	"go-gin/internal/components/logx"
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				m := map[string]interface{}{
					"error": err,
					"file":  utils.FileWithLineNum(),
				}
				logx.WithContext(ctx).Error("panic", m)
				httpx.Error(ctx, consts.ErrInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()

	}
}

package middlewares

import (
	"go-gin/internal/components/logx"
	"go-gin/internal/utils"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logx.PanicLoggerInstance.Fatal().
					Ctx(ctx).
					Any("error", err).
					Str("file", utils.FileWithLineNum()).
					Send()
				ctx.Abort()
			}
		}()
		ctx.Next()

	}
}

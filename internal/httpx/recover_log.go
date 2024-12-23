package httpx

import (
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/errorx"
	"go-gin/internal/util"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				m := map[string]any{
					"error": fmt.Sprintf("%v", r),
					"file":  util.FileWithLineNum(),
				}
				logx.WithContext(ctx).Error("panic", m)
				Error(NewContext(ctx), errorx.ErrInternalServerError)
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}

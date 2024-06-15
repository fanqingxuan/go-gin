package middlewares

import (
	"fmt"
	"go-gin/consts"
	"go-gin/utils/httpx"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println(string(debug.Stack()))
				httpx.Error(ctx, consts.ErrInternalServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

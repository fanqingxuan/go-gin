package middlewares

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/utils/httpx"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println(string(debug.Stack()))
				httpx.Error(ctx, errorx.New(http.StatusInternalServerError, "服务器内部错误"))
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

package middlewares

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/errorx"
	"go-gin/internal/ginx/httpx"

	"github.com/gin-gonic/gin"
)

func dbCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if db.IsNotOpened() {
			err := db.Connect()
			if err != nil {
				logx.WithContext(ctx).Error("connect db again", err.Error())
				httpx.Error(ctx, errorx.NewDBError(err))
				ctx.Abort()
			}
		}
		ctx.Next()

	}
}

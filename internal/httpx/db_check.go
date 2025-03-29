package httpx

import (
	"go-gin/internal/component/db"
	"go-gin/internal/component/logx"
	"go-gin/internal/errorx"

	"github.com/gin-gonic/gin"
)

func dbCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !db.IsConnected() {
			err := db.Connect()
			if err != nil {
				logx.WithContext(ctx).Error("connect db", err.Error())
				Error(NewContext(ctx), errorx.ErrInternalServerError)
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}

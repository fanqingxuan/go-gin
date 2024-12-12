package middlewares

import (
	"go-gin/internal/components/db"
	"go-gin/internal/components/logx"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
)

func dbCheck() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (interface{}, error) {
		if !db.IsConnected() {
			err := db.Connect()
			if err != nil {
				logx.WithContext(ctx).Error("connect db again", err.Error())
				return nil, errorx.TryToDBError(err)
			}
		}
		return nil, nil
	}
}

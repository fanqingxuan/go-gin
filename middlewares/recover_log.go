package middlewares

import (
	"go-gin/internal/httpx"
)

func recoverLog() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (interface{}, error) {
		// defer func() {
		// 	if err := recover(); err != nil {
		// 		m := map[string]interface{}{
		// 			"error": fmt.Sprintf("%v", err),
		// 			"file":  utils.FileWithLineNum(),
		// 		}
		// 		logx.WithContext(ctx).Error("panic", m)
		// 		httpx.Error(ctx, errorx.ErrInternalServerError)
		// 		ctx.Abort()
		// 	}
		// }()
		ctx.Next()
		return nil, nil
	}
}

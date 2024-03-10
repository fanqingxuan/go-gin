package middlewares

import "github.com/gin-gonic/gin"

func RegisterGlobalMiddlewares(r *gin.Engine) {

	r.Use(traceId())
	r.Use(requestLog())
	r.Use(recoverLog())
}

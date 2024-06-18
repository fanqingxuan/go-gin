package middlewares

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.Use(traceId())
	r.Use(requestLog())
	r.Use(dbCheck())
	r.Use(recoverLog())
}

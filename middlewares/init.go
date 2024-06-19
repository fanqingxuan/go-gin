package middlewares

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.Use(dbCheck())
	r.Use(recoverLog())
}

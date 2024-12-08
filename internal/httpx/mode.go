package httpx

import "github.com/gin-gonic/gin"

func SetDebugMode() {
	gin.SetMode(gin.DebugMode)
}

func SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

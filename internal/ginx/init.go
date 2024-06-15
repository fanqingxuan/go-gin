package ginx

import (
	"go-gin/internal/ginx/validators"

	"github.com/gin-gonic/gin"
)

func Init(isDebug bool) *gin.Engine {
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.HandleMethodNotAllowed = true
	validators.Init()
	return engine
}

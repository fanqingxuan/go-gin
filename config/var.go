package config

import (
	"go-gin/internal/ginx/httpx"
	"go-gin/internal/traceid"
)

func InitGlobalVars() {
	// httpx.DefaultSuccessCodeValue = 0
	httpx.DefaultSuccessMessageValue = "成功"

	traceid.TraceIdFieldName = "requestId"
}

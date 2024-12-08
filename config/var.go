package config

import (
	"go-gin/internal/httpx"
	"go-gin/internal/traceid"
)

func InitGlobalVars() {
	httpx.DefaultSuccessCodeValue = 200
	httpx.CodeFieldName = "status"
	httpx.ResultFieldName = "data"
	httpx.MessageFieldName = "msg"
	httpx.DefaultSuccessMessageValue = "成功"
	traceid.TraceIdFieldName = "requestId"
}

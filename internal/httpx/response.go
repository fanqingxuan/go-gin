package httpx

import (
	"context"
	"go-gin/internal/environment"
	"go-gin/internal/errorx"
	"go-gin/internal/traceid"
	"net/http"
)

var (
	CodeFieldName    = "code"
	ResultFieldName  = "data"
	MessageFieldName = "message"
)

var DefaultSuccessCodeValue = http.StatusOK
var DefaultSuccessMessageValue = "操作成功"

type Result struct {
	Code    int
	Message string
	Data    any
	TraceId string
}

func Ok(ctx *Context, data any) {
	result := Result{
		Code:    DefaultSuccessCodeValue,
		Data:    data,
		Message: DefaultSuccessMessageValue,
	}
	ctx.JSON(http.StatusOK, transform(ctx, result))
}

func Error(ctx *Context, err error) {
	var httpStatus int
	var code int
	var message string

	switch e := err.(type) {
	case errorx.ServerError:
		code = e.Code
		httpStatus = e.Code
		message = e.Msg
	case errorx.BizError:
		message = e.Msg
		httpStatus = http.StatusOK
		code = e.Code
	case errorx.RedisError:
		if environment.IsDebugMode() {
			message = e.Error()
		} else {
			message = "服务器内部错误"
		}
		httpStatus = http.StatusInternalServerError
		code = errorx.ErrCodeRedisOperateFailed
	case errorx.DBError:
		if environment.IsDebugMode() {
			message = e.Error()
		} else {
			message = "服务器内部错误"
		}
		httpStatus = http.StatusInternalServerError
		code = errorx.ErrCodeDBOperateFailed
	case error:
		httpStatus = http.StatusOK
		code = errorx.ErrCodeDefault
		message = err.Error()
	}
	result := Result{
		Code:    code,
		Message: message,
	}
	ctx.JSON(httpStatus, transform(ctx, result))
}

func Handle(ctx *Context, data any, err error) {
	if err != nil {
		Error(ctx, err)
	} else {
		Ok(ctx, data)
	}
}

func transform(ctx context.Context, result Result) map[string]any {
	s, _ := ctx.Value(traceid.TraceIdFieldName).(string)

	return map[string]any{
		CodeFieldName:            result.Code,
		MessageFieldName:         result.Message,
		ResultFieldName:          result.Data,
		traceid.TraceIdFieldName: s,
	}
}

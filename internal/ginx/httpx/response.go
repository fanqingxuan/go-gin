package httpx

import (
	"context"
	"go-gin/internal/errorx"
	"go-gin/internal/traceid"
	"net/http"

	"github.com/gin-gonic/gin"
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

func Ok(ctx *gin.Context, data any) {
	OkWithMessage(ctx, data, DefaultSuccessMessageValue)
}

func OkResponse(ctx *gin.Context) {
	OkWithMessage(ctx, nil, DefaultSuccessMessageValue)
}

func OkWithMessage(ctx *gin.Context, data any, msg string) {
	result := Result{
		Code:    DefaultSuccessCodeValue,
		Data:    data,
		Message: msg,
	}
	ctx.JSON(http.StatusOK, transform(ctx, result))
}

func Error(ctx *gin.Context, err error) {
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
		message = "服务器内部错误"
		httpStatus = http.StatusInternalServerError
		code = errorx.ErrCodeRedisOperateFailed
	case errorx.DBError:
		message = "服务器内部错误"
		httpStatus = http.StatusInternalServerError
		code = errorx.ErrCodeDBOperateFailed
	case error:
		httpStatus = http.StatusOK
		code = errorx.ErrCodeDefaultCommon
		message = err.Error()
	}
	result := Result{
		Code:    code,
		Message: message,
	}
	ctx.JSON(httpStatus, transform(ctx, result))
}

func Handle(ctx *gin.Context, data any, err error) {
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

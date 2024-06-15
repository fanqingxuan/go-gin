package httpx

import (
	"context"
	"go-gin/utils/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CodeFieldName = "code"
var ResultFieldName = "data"
var MessageFieldName = "message"
var RequestIdFieldName = "requestId"

var DefaultSuccessCodeValue = http.StatusOK

type Result struct {
	Code      int
	Message   string
	Data      any
	RequestId string
}

func Ok(ctx *gin.Context, data any) {
	result := Result{
		Code:    DefaultSuccessCodeValue,
		Data:    data,
		Message: "操作成功",
	}
	ctx.JSON(http.StatusOK, transform(ctx, result))
}

func Error(ctx *gin.Context, err error) {
	var httpStatus int
	var code int
	var message string

	switch e := err.(type) {
	case errorx.HttpError:
		code = e.Code
		httpStatus = e.Code
		message = http.StatusText(httpStatus)
	case errorx.BizError:
		message = e.Msg
		httpStatus = http.StatusOK
		code = DefaultSuccessCodeValue
	case error:
		message = "服务器内部错误"
		httpStatus = http.StatusInternalServerError
		code = httpStatus
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
	s, _ := ctx.Value("requestId").(string)

	return map[string]any{
		CodeFieldName:      result.Code,
		MessageFieldName:   result.Message,
		ResultFieldName:    result.Data,
		RequestIdFieldName: s,
	}
}

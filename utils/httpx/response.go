package httpx

import (
	"context"
	"fmt"
	"go-gin/internal/errorx"
	"net/http"

	"github.com/gin-gonic/gin"
)

var CodeFieldName = "status"
var ResultFieldName = "data"
var MessageFieldName = "message"
var RequestIdFieldName = "requestId"

func Ok(ctx *gin.Context, data any) {
	OkWithMessage(ctx, data, "操作成功")
}

func OkWithMessage(ctx *gin.Context, data any, message string) {
	result := Result{
		Code:    http.StatusOK,
		Data:    data,
		Message: message,
	}
	ctx.JSON(http.StatusOK, transform(ctx, result))
}

func Error(ctx *gin.Context, err error) {
	httpStatus := ctx.Writer.Status()
	code := httpStatus
	message := http.StatusText(httpStatus)

	switch e := err.(type) {
	case errorx.MYError:
		message = e.Msg
		httpStatus = http.StatusOK
		code = e.Code
	case error:
		message = "服务器内部错误"
		httpStatus = http.StatusInternalServerError
		code = httpStatus
		fmt.Println(err)
	}
	result := Result{
		Code:    code,
		Message: message,
	}
	ctx.JSON(httpStatus, transform(ctx, result))
}

func transform(ctx context.Context, result Result) map[string]any {
	s, _ := ctx.Value("requestId").(string)

	return map[string]any{
		CodeFieldName:      result.Code,
		MessageFieldName:   result.Message,
		ResultFieldName:    result.Data,
		RequestIdFieldName: fmt.Sprintf("%s", s),
	}
}

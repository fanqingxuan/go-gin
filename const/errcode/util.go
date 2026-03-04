package errcode

import (
	"go-gin/internal/errorx"
)

// NewDefault 创建默认业务错误
func NewDefault(msg string) errorx.BizError {
	return errorx.BizError{Code: errorx.ErrCodeBizDefault, Msg: msg}
}

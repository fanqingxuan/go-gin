package errcode

import (
	"errors"
	"go-gin/internal/errorx"
)

func New(code int, msg string) errorx.BizError {
	return errorx.BizError{Code: code, Msg: msg}
}

func NewDefault(msg string) errorx.BizError {
	return NewError(errors.New(msg))
}

func NewError(err error) errorx.BizError {
	return errorx.BizError{Code: errorx.ErrCodeBizDefault, Msg: err.Error()}
}

func IsRecordNotFound(err error) bool {
	return errorx.IsRecordNotFound(err)
}

func IsError(err error) bool {
	return errorx.IsError(err)
}

package errcode

import (
	"go-gin/internal/errorx"
)

var (
	ErrXXX       = errorx.New(number, message)
)

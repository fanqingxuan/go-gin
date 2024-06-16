package consts

import (
	"go-gin/internal/errorx"
	"net/http"
)

const (
	ErrCodeUserNotFound = 2001
)

var (
	ErrMethodNotAllowed    = errorx.NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = errorx.NewServerError(http.StatusNotFound)
	ErrInternalServerError = errorx.NewServerError(http.StatusInternalServerError)

	ErrUserNotFound = errorx.New(ErrCodeUserNotFound, "用户不存在")
)

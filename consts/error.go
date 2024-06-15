package consts

import (
	"go-gin/internal/errorx"
	"net/http"
)

var (
	ErrMethodNotAllowed    = errorx.NewHHttpError(http.StatusMethodNotAllowed)
	ErrNoRoute             = errorx.NewHHttpError(http.StatusNotFound)
	ErrInternalServerError = errorx.NewHHttpError(http.StatusInternalServerError)

	ErrUserNotFound = errorx.New(2001, "用户不存在")
)

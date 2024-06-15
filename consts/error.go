package consts

import (
	"go-gin/internal/errorx"
	"net/http"
)

var (
	ErrMethodNotAllowed    = errorx.NewHHttpError(http.StatusMethodNotAllowed)
	ErrNoRoute             = errorx.NewHHttpError(http.StatusNotFound)
	ErrInternalServerError = errorx.NewHHttpError(http.StatusInternalServerError)

	ErrUserNotFound = errorx.New(1002, "用户不存在")
)

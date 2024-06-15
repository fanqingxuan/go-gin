package consts

import (
	"go-gin/utils/errorx"
	"net/http"
)

var (
	ErrMethodNotAllowed    = errorx.NewHHttpError(http.StatusMethodNotAllowed, "方法不允许")
	ErrNoRoute             = errorx.NewHHttpError(http.StatusNotFound, "路由不存在")
	ErrInternalServerError = errorx.NewHHttpError(http.StatusInternalServerError, "服务器内部错误")

	ErrUserNotFound = errorx.New(1002, "用户不存在")
)

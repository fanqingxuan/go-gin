package consts

import (
	"go-gin/internal/errorx"
	"net/http"
)

const (
	ErrCodeUserNotFound = 2001
)

var (
	// http错误
	ErrMethodNotAllowed    = errorx.NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = errorx.NewServerError(http.StatusNotFound)
	ErrInternalServerError = errorx.NewServerError(http.StatusInternalServerError)

	// 错误是5位数的数字

	// 1开头的是用于定义数据库、redis、文件解析、json解析等非业务错误
	// 数据库错误
	ErrDBConnectFailed      = errorx.New(10001, "数据库连接异常")
	ErrDBCreateRecordFailed = errorx.New(10001, "创建记录失败")
	ErrDBDeleteRecordFailed = errorx.New(10002, "删除记录失败")
	ErrDBModifyRecordFailed = errorx.New(10003, "修改记录失败")
	ErrDBQueryRecordFailed  = errorx.New(10004, "查询记录失败")

	// 以下定义业务上的错误
	ErrUserNotFound = errorx.New(ErrCodeUserNotFound, "用户不存在")
)

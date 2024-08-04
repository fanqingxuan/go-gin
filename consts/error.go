package consts

import (
	"go-gin/internal/errorx"
	"net/http"
)

var (
	// http错误
	ErrMethodNotAllowed    = errorx.NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = errorx.NewServerError(http.StatusNotFound)
	ErrInternalServerError = errorx.NewServerError(http.StatusInternalServerError)

	// 错误是5位数的数字

	// 1开头的是用于定义数据库、redis、文件解析、json解析等非业务错误
	// 数据库错误
	ErrDBConnectFailed      = errorx.New(11000, "数据库连接异常")
	ErrDBCreateRecordFailed = errorx.New(11001, "创建记录失败")
	ErrDBDeleteRecordFailed = errorx.New(11002, "删除记录失败")
	ErrDBModifyRecordFailed = errorx.New(11003, "修改记录失败")
	ErrDBQueryRecordFailed  = errorx.New(11004, "查询记录失败")

	ErrThirdPartyAPIRequestFailed = errorx.New(11005, "第三方接口请求失败")

	// 以下定义业务上的错误
	ErrUserNotFound       = errorx.New(20001, "用户不存在")
	ErrUserNameOrPwdFaild = errorx.New(20002, "用户名或者密码错误")
	ErrUserMustLogin      = errorx.New(20003, "请先登录")
)

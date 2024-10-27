package consts

import (
	"go-gin/internal/errorx"
)

var (
	// 以下定义业务上的错误,注意1开头的是系统错误
	ErrUserNotFound       = errorx.New(20001, "用户不存在")
	ErrUserNameOrPwdFaild = errorx.New(20002, "用户名或者密码错误")
	ErrUserMustLogin      = errorx.New(20003, "请先登录")
	ErrUserNeedLoginAgain = errorx.New(20004, "token已过期,请重新登录")
)

package enum

import (
	"go-gin/internal/etype"
)

// UserStatus 用户状态
type UserStatus struct {
	etype.BaseEnum
}

// 用户状态常量
var (
	USER_STATUS_NORMAL   = etype.NewEnum[UserStatus](1, "正常")
	USER_STATUS_DISABLED = etype.NewEnum[UserStatus](2, "禁用")
	USER_STATUS_DELETED  = etype.NewEnum[UserStatus](3, "已删除")
)

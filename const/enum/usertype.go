package enum

import (
	"go-gin/internal/etype"
)

// UserType 用户类型
type UserType struct {
	etype.BaseEnum
}

// 用户类型常量
var (
	USER_TYPE_NORMAL     = etype.NewEnum[UserType](1, "正常添加")
	USER_TYPE_FROM_THIRD = etype.NewEnum[UserType](2, "从第三方导入")
	USER_TYPE_SUPPLIER   = etype.NewEnum[UserType](3, "供应商用户")
)

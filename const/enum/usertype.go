package enum

import (
	"go-gin/internal/etype"
)

// UserType 用户类型
type UserType struct {
	etype.BaseEnum
}

// 定义用户类型常量
var (
	UserTypeNormal    = etype.NewEnum[UserType](1, "正常添加")
	UserTypeFromThird = etype.NewEnum[UserType](2, "从第三方导入")
	UserTypeSupplier  = etype.NewEnum[UserType](3, "供应商用户")
)

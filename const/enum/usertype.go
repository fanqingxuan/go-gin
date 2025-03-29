package enum

import (
	"fmt"
	"go-gin/internal/etype"
)

// 用户类型
type UserType etype.NumEnum

const (
	// 正常添加
	UserTypeNormal UserType = 1
	// 从第三方导入
	UserTypeFromThird = 2
	// 供应商用户
	UserTypeSupplier = 3
)

var userTypeMap = map[UserType]string{
	UserTypeNormal:    "正常添加",
	UserTypeFromThird: "从第三方导入",
	UserTypeSupplier:  "供应商用户",
}

func (b UserType) String() string {
	return etype.String(b, userTypeMap)
}

func (b UserType) Equal(other UserType) bool {
	return etype.Equal(b, other)
}

func (b UserType) Format(f fmt.State, r rune) {
	etype.Format(f, b)
}

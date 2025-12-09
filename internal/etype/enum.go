package etype

import (
	"database/sql/driver"
	"fmt"
)

// CodeGetter 用于枚举比较的接口
type CodeGetter interface {
	Code() int
}

// BaseEnum 基础枚举结构体
type BaseEnum struct {
	code   int
	desc   string
	prefix PrefixType
}

// NewBaseEnum 创建基础枚举
func NewBaseEnum(prefix PrefixType, code int, desc string) *BaseEnum {
	return &BaseEnum{
		code:   code,
		desc:   desc,
		prefix: prefix,
	}
}

// Code 获取状态码
func (e *BaseEnum) Code() int {
	return e.code
}

// Desc 获取描述
func (e *BaseEnum) Desc() string {
	return e.desc
}

// String 实现 Stringer 接口
func (e *BaseEnum) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s(%d)", e.desc, e.code)
}

// Equal 比较两个枚举是否相等
func (e *BaseEnum) Equal(other CodeGetter) bool {
	if e == nil && other == nil {
		return true
	}
	if e == nil || other == nil {
		return false
	}
	return e.code == other.Code()
}

// Value 实现 driver.Valuer 接口（被生成代码调用）
func (e *BaseEnum) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}
	return int64(e.code), nil
}

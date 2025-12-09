package etype

import (
	"database/sql/driver"
	"encoding/json"
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
	base := &BaseEnum{
		code:   code,
		desc:   desc,
		prefix: prefix,
	}
	return base
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

// Value 实现 driver.Valuer 接口
func (e *BaseEnum) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}
	return int64(e.code), nil
}

// MarshalJSON 实现 json.Marshaler 接口
func (e *BaseEnum) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}
	return json.Marshal(e.code)
}

// Scan 实现 sql.Scanner 接口
func (s *BaseEnum) Scan(value any) error {
	if value == nil {
		*s = BaseEnum{}
		return nil
	}

	var code int
	switch v := value.(type) {
	case int64:
		code = int(v)
	case int:
		code = v
	default:
		return fmt.Errorf("不支持的类型转换: %T", value)
	}
	m := getAll(s.prefix)
	if base, ok := m[code]; ok {
		s.code = code
		s.desc = base.Desc()
		return nil
	}
	return fmt.Errorf("未知的code码: %d", code)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *BaseEnum) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*s = BaseEnum{}
		return nil
	}

	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	m := getAll(s.prefix)
	if base, ok := m[code]; ok {
		s.code = code
		s.desc = base.Desc()
		return nil
	}
	return fmt.Errorf("未知的code码: %d", code)
}

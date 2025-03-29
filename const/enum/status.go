package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Status 表示一个状态
type Status struct {
	code int
	desc string
}

var _ IEnum = (*Status)(nil)

// 状态码和描述的映射关系
var statusMap = make(map[int]*Status)

// NewStatus 创建一个新的状态
func NewStatus(code int, desc string) *Status {
	// 自动维护映射关系
	status := &Status{
		code: code,
		desc: desc,
	}
	statusMap[code] = status
	return status
}

// Status获取结构体
func ParseStatus(code int) (*Status, error) {
	if status, ok := statusMap[code]; ok {
		// 创建一个新的 Status 副本
		return &Status{
			code: status.code,
			desc: status.desc,
		}, nil
	}
	return nil, fmt.Errorf("未知状态码: %d", code)
}

// Code 获取状态码
func (s *Status) Code() int {
	return s.code
}

// Desc 获取描述
func (s *Status) Desc() string {
	return s.desc
}

// String 实现 Stringer 接口
func (s *Status) String() string {
	return fmt.Sprintf("%s(%d)", s.desc, s.code)
}

// Equal 比较两个 Status 是否相等
func (s *Status) Equal(other *Status) bool {
	if s == nil || other == nil {
		return s == other
	}
	return s.code == other.code && s.desc == other.desc
}

// Scan 实现 sql.Scanner 接口
func (s *Status) Scan(value interface{}) error {
	if value == nil {
		s = nil
		return nil
	}

	switch v := value.(type) {
	case int64:
		s.code = int(v)
	case int:
		s.code = v
	default:
		return fmt.Errorf("不支持的类型转换: %T", value)
	}

	// 从映射中获取状态
	if status, ok := statusMap[s.code]; ok {
		s.desc = status.desc
	} else {
		return fmt.Errorf("未知状态码(%v)", s.code)
	}
	return nil
}

// Value 实现 driver.Valuer 接口
func (s *Status) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return int64(s.code), nil
}

// MarshalJSON 实现 json.Marshaler 接口
func (s *Status) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	return json.Marshal(s.code)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *Status) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		s = nil
		return nil
	}

	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}

	// 验证状态码是否有效
	if status, ok := statusMap[code]; ok {
		s.code = status.code
		s.desc = status.desc
		return nil
	}
	return fmt.Errorf("UnmarshalJSON出现错误,未定义的状态码: %d", code)
}

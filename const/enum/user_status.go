package enum

import (
	"fmt"
)

const (
	// PrefixUserStatus 用户状态前缀
	PrefixUserStatus PrefixType = "user_status"
)

// UserStatus 用户状态
type UserStatus struct {
	BaseEnum
}

// 定义用户状态常量
var (
	USER_STATUS_NORMAL   = NewUserStatus(1, "正常")
	USER_STATUS_DISABLED = NewUserStatus(2, "禁用")
	USER_STATUS_DELETED  = NewUserStatus(3, "已删除")
)

// NewUserStatus 创建用户状态
func NewUserStatus(code int, desc string) *UserStatus {
	status := &UserStatus{
		BaseEnum: BaseEnum{
			code: code,
			desc: desc,
		},
	}
	Set(PrefixUserStatus, code, desc)
	return status
}

// ParseUserStatus 解析用户状态
func ParseUserStatus(code int) (*UserStatus, error) {
	// 使用EnumMapManager获取描述
	if desc, ok := Get(PrefixUserStatus, code); ok {
		return &UserStatus{
			BaseEnum: BaseEnum{
				code: code,
				desc: desc,
			},
		}, nil
	}
	return nil, fmt.Errorf("未知的用户状态码: %d", code)
}

// Scan 实现 sql.Scanner 接口
func (s *UserStatus) Scan(value interface{}) error {
	return s.BaseEnum.Scan(value, GetAll(PrefixUserStatus))
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *UserStatus) UnmarshalJSON(data []byte) error {
	return s.BaseEnum.UnmarshalJSON(data, GetAll(PrefixUserStatus))

}

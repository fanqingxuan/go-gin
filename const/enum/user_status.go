package enum

import (
	"go-gin/internal/etype"
)

const PrefixUserStatus etype.PrefixType = "user_status"

// 定义用户状态常量
var (
	USER_STATUS_NORMAL   = NewUserStatus(1, "正常")
	USER_STATUS_DISABLED = NewUserStatus(2, "禁用")
	USER_STATUS_DELETED  = NewUserStatus(3, "已删除")
)

// UserStatus 用户状态
type UserStatus struct {
	etype.BaseEnum
}

// NewUserStatus 创建用户状态
func NewUserStatus(code int, desc string) *UserStatus {
	return &UserStatus{
		BaseEnum: etype.CreateBaseEnumAndSetMap(PrefixUserStatus, code, desc),
	}
}

// ParseUserStatus 解析用户状态
func ParseUserStatus(code int) (*UserStatus, error) {

	base, err := etype.ParseBaseEnum(PrefixUserStatus, code)
	if err != nil {
		return nil, err
	}

	return &UserStatus{BaseEnum: base}, nil
}

// Scan 实现 sql.Scanner 接口
func (s *UserStatus) Scan(value interface{}) error {
	return s.BaseEnum.Scan(value, PrefixUserStatus)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *UserStatus) UnmarshalJSON(data []byte) error {
	return s.BaseEnum.UnmarshalJSON(data, PrefixUserStatus)

}

package etype

import "fmt"

func CreateBaseEnumAndSetMap(prefix PrefixType, code int, desc string) BaseEnum {
	baseenum := NewBaseEnum(code, desc)
	Set(prefix, baseenum)
	return *baseenum
}

// ParseBaseEnum 通用的枚举解析方法
func ParseBaseEnum(prefix PrefixType, code int) (BaseEnum, error) {
	if base, ok := Get(prefix, code); ok {
		return CreateBaseEnumAndSetMap(prefix, code, base.Desc()), nil
	}
	return BaseEnum{}, fmt.Errorf("未知的enum码: %d", code)
}

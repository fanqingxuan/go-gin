package etype

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// PrefixType 前缀类型（用于区分不同枚举类型）
type PrefixType string

// ValueType 枚举值映射
type ValueType map[int]*BaseEnum

// 全局枚举注册表
var (
	enumMap      = make(map[PrefixType]ValueType)
	enumMapMutex sync.RWMutex
)

// set 注册枚举值
func set(prefix PrefixType, enum *BaseEnum) {
	enumMapMutex.Lock()
	defer enumMapMutex.Unlock()

	if _, ok := enumMap[prefix]; !ok {
		enumMap[prefix] = make(ValueType)
	}
	enumMap[prefix][enum.code] = enum
}

// get 获取枚举值
func get(prefix PrefixType, code int) (*BaseEnum, bool) {
	enumMapMutex.RLock()
	defer enumMapMutex.RUnlock()

	if prefixMap, ok := enumMap[prefix]; ok {
		value, exists := prefixMap[code]
		return value, exists
	}
	return nil, false
}

// NewEnum 泛型枚举构造函数
// T 必须是嵌入了 BaseEnum 的结构体
func NewEnum[T any](code int, desc string) *T {
	var t T
	prefix := getPrefixFromType[T]()

	// 创建并注册 BaseEnum
	base := NewBaseEnum(prefix, code, desc)
	set(prefix, base)

	// 使用反射设置 BaseEnum 字段
	v := reflect.ValueOf(&t).Elem()
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Type() == reflect.TypeOf(BaseEnum{}) && field.CanSet() {
				field.Set(reflect.ValueOf(*base))
				break
			}
		}
	}
	return &t
}

// Parse 通过 code 获取枚举（被 ScanEnum/UnmarshalEnum 调用）
func Parse[T any](code int) (*T, error) {
	prefix := getPrefixFromType[T]()
	base, ok := get(prefix, code)
	if !ok {
		return nil, fmt.Errorf("未知的枚举值: %d", code)
	}

	var t T
	v := reflect.ValueOf(&t).Elem()
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Type() == reflect.TypeOf(BaseEnum{}) && field.CanSet() {
				field.Set(reflect.ValueOf(*base))
				break
			}
		}
	}
	return &t, nil
}

// ScanEnum 数据库扫描（被生成代码调用）
func ScanEnum[T any](value any) (*T, error) {
	if value == nil {
		return nil, nil
	}

	var code int
	switch v := value.(type) {
	case int64:
		code = int(v)
	case int:
		code = v
	default:
		return nil, fmt.Errorf("不支持的类型: %T", value)
	}

	return Parse[T](code)
}

// UnmarshalEnum JSON 反序列化（被生成代码调用）
func UnmarshalEnum[T any](data []byte) (*T, error) {
	if len(data) == 0 || string(data) == "null" {
		return nil, nil
	}

	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return nil, err
	}

	return Parse[T](code)
}

// getPrefixFromType 从类型名生成 prefix
// 例如: UserStatus -> user_status
func getPrefixFromType[T any]() PrefixType {
	var t T
	typeName := reflect.TypeOf(t).Name()
	return PrefixType(toSnakeCase(typeName))
}

// toSnakeCase 驼峰转下划线
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

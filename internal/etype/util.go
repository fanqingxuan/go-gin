package etype

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// PrefixType 前缀类型
type PrefixType string
type ValueType map[int]*BaseEnum

// 包级别的二维map变量
var (
	enumMap      = make(map[PrefixType]ValueType)
	enumMapMutex sync.RWMutex
)

// set 设置枚举值
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

// getAll 获取指定前缀的所有值
func getAll(prefix PrefixType) ValueType {
	enumMapMutex.RLock()
	defer enumMapMutex.RUnlock()

	if prefixMap, ok := enumMap[prefix]; ok {
		result := make(ValueType, len(prefixMap))
		for k, v := range prefixMap {
			result[k] = v
		}
		return result
	}
	return make(ValueType)
}

func createBaseEnumAndSetMap(prefix PrefixType, code int, desc string) BaseEnum {
	baseenum := NewBaseEnum(prefix, code, desc)
	set(prefix, baseenum)
	return *baseenum
}

// NewEnum 泛型枚举构造函数，自动推断 prefix
// T 必须是嵌入了 BaseEnum 的结构体
func NewEnum[T any](code int, desc string) *T {
	var t T
	prefix := getPrefixFromType[T]()
	base := createBaseEnumAndSetMap(prefix, code, desc)

	// 使用反射设置 BaseEnum 字段
	v := reflect.ValueOf(&t).Elem()
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Type() == reflect.TypeOf(BaseEnum{}) && field.CanSet() {
				field.Set(reflect.ValueOf(base))
				break
			}
		}
	}
	return &t
}

// Parse 泛型解析函数，通过 code 获取枚举
func Parse[T any](code int) (*T, error) {
	prefix := getPrefixFromType[T]()
	base, ok := get(prefix, code)
	if !ok {
		return nil, fmt.Errorf("未知的enum码: %d", code)
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

// getPrefixFromType 从类型名生成 prefix
// 例如: OrderStatus -> order_status
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

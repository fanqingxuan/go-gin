package etype

import (
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

// Set 设置枚举值
func Set(prefix PrefixType, enum *BaseEnum) {

	enumMapMutex.Lock()
	defer enumMapMutex.Unlock()

	if _, ok := enumMap[prefix]; !ok {
		enumMap[prefix] = make(ValueType)
	}
	enumMap[prefix][enum.code] = enum
}

// Get 获取枚举值
func Get(prefix PrefixType, code int) (*BaseEnum, bool) {

	enumMapMutex.RLock()
	defer enumMapMutex.RUnlock()

	if prefixMap, ok := enumMap[prefix]; ok {
		value, exists := prefixMap[code]
		return value, exists
	}
	return &BaseEnum{}, false
}

// GetAll 获取指定前缀的所有值
func GetAll(prefix PrefixType) ValueType {
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

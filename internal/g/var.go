package g

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Var 通用变量类型，参考 GoFrame gvar
type Var struct {
	value any
}

// NewVar 创建 Var
func NewVar(value any) *Var {
	return &Var{value: value}
}

// Val 返回原始值
func (v *Var) Val() any {
	if v == nil {
		return nil
	}
	return v.value
}

// Interface 是 Val 的别名
func (v *Var) Interface() any {
	return v.Val()
}

// Set 设置值
func (v *Var) Set(value any) {
	v.value = value
}

// IsNil 判断是否为 nil
func (v *Var) IsNil() bool {
	if v == nil {
		return true
	}
	return v.value == nil
}

// IsEmpty 判断是否为空
func (v *Var) IsEmpty() bool {
	if v.IsNil() {
		return true
	}
	switch val := v.value.(type) {
	case string:
		return val == ""
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(val).Int() == 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(val).Uint() == 0
	case float32, float64:
		return reflect.ValueOf(val).Float() == 0
	case bool:
		return !val
	case []byte:
		return len(val) == 0
	default:
		rv := reflect.ValueOf(val)
		switch rv.Kind() {
		case reflect.Slice, reflect.Map, reflect.Array:
			return rv.Len() == 0
		}
	}
	return false
}

// String 转换为 string
func (v *Var) String() string {
	if v.IsNil() {
		return ""
	}
	switch val := v.value.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(val).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(val).Uint(), 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case time.Time:
		return val.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%v", val)
	}
}

// Bytes 转换为 []byte
func (v *Var) Bytes() []byte {
	if v.IsNil() {
		return nil
	}
	if b, ok := v.value.([]byte); ok {
		return b
	}
	return []byte(v.String())
}

// Bool 转换为 bool
func (v *Var) Bool() bool {
	if v.IsNil() {
		return false
	}
	switch val := v.value.(type) {
	case bool:
		return val
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(val).Int() != 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(val).Uint() != 0
	case float32, float64:
		return reflect.ValueOf(val).Float() != 0
	case string:
		return val != "" && val != "0" && val != "false"
	case []byte:
		return len(val) > 0
	}
	return true
}

// Int 转换为 int
func (v *Var) Int() int {
	return int(v.Int64())
}

// Int8 转换为 int8
func (v *Var) Int8() int8 {
	return int8(v.Int64())
}

// Int16 转换为 int16
func (v *Var) Int16() int16 {
	return int16(v.Int64())
}

// Int32 转换为 int32
func (v *Var) Int32() int32 {
	return int32(v.Int64())
}

// Int64 转换为 int64
func (v *Var) Int64() int64 {
	if v.IsNil() {
		return 0
	}
	switch val := v.value.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	case []byte:
		i, _ := strconv.ParseInt(string(val), 10, 64)
		return i
	}
	return 0
}

// Uint 转换为 uint
func (v *Var) Uint() uint {
	return uint(v.Uint64())
}

// Uint8 转换为 uint8
func (v *Var) Uint8() uint8 {
	return uint8(v.Uint64())
}

// Uint16 转换为 uint16
func (v *Var) Uint16() uint16 {
	return uint16(v.Uint64())
}

// Uint32 转换为 uint32
func (v *Var) Uint32() uint32 {
	return uint32(v.Uint64())
}

// Uint64 转换为 uint64
func (v *Var) Uint64() uint64 {
	if v.IsNil() {
		return 0
	}
	switch val := v.value.(type) {
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return val
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		i, _ := strconv.ParseUint(val, 10, 64)
		return i
	case []byte:
		i, _ := strconv.ParseUint(string(val), 10, 64)
		return i
	}
	return 0
}

// Float32 转换为 float32
func (v *Var) Float32() float32 {
	return float32(v.Float64())
}

// Float64 转换为 float64
func (v *Var) Float64() float64 {
	if v.IsNil() {
		return 0
	}
	switch val := v.value.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	case []byte:
		f, _ := strconv.ParseFloat(string(val), 64)
		return f
	}
	return 0
}

// Time 转换为 time.Time
func (v *Var) Time(format ...string) time.Time {
	if v.IsNil() {
		return time.Time{}
	}
	switch val := v.value.(type) {
	case time.Time:
		return val
	case *time.Time:
		if val != nil {
			return *val
		}
		return time.Time{}
	case string:
		layout := "2006-01-02 15:04:05"
		if len(format) > 0 {
			layout = format[0]
		}
		t, _ := time.ParseInLocation(layout, val, time.Local)
		return t
	case []byte:
		layout := "2006-01-02 15:04:05"
		if len(format) > 0 {
			layout = format[0]
		}
		t, _ := time.ParseInLocation(layout, string(val), time.Local)
		return t
	case int64:
		return time.Unix(val, 0)
	}
	return time.Time{}
}

// Duration 转换为 time.Duration
func (v *Var) Duration() time.Duration {
	if v.IsNil() {
		return 0
	}
	switch val := v.value.(type) {
	case time.Duration:
		return val
	case int64:
		return time.Duration(val)
	case string:
		d, _ := time.ParseDuration(val)
		return d
	}
	return time.Duration(v.Int64())
}

// Ints 转换为 []int
func (v *Var) Ints() []int {
	arr := v.Interfaces()
	result := make([]int, len(arr))
	for i, item := range arr {
		result[i] = NewVar(item).Int()
	}
	return result
}

// Int64s 转换为 []int64
func (v *Var) Int64s() []int64 {
	arr := v.Interfaces()
	result := make([]int64, len(arr))
	for i, item := range arr {
		result[i] = NewVar(item).Int64()
	}
	return result
}

// Strings 转换为 []string
func (v *Var) Strings() []string {
	arr := v.Interfaces()
	result := make([]string, len(arr))
	for i, item := range arr {
		result[i] = NewVar(item).String()
	}
	return result
}

// Floats 转换为 []float64
func (v *Var) Floats() []float64 {
	arr := v.Interfaces()
	result := make([]float64, len(arr))
	for i, item := range arr {
		result[i] = NewVar(item).Float64()
	}
	return result
}

// Interfaces 转换为 []any
func (v *Var) Interfaces() []any {
	if v.IsNil() {
		return nil
	}
	switch val := v.value.(type) {
	case []any:
		return val
	default:
		rv := reflect.ValueOf(val)
		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			result := make([]any, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				result[i] = rv.Index(i).Interface()
			}
			return result
		}
	}
	return []any{v.value}
}

// Slice 是 Interfaces 的别名
func (v *Var) Slice() []any {
	return v.Interfaces()
}

// Array 是 Interfaces 的别名
func (v *Var) Array() []any {
	return v.Interfaces()
}

// Map 转换为 map[string]any
func (v *Var) Map() map[string]any {
	if v.IsNil() {
		return nil
	}
	if m, ok := v.value.(map[string]any); ok {
		return m
	}
	rv := reflect.ValueOf(v.value)
	if rv.Kind() == reflect.Map {
		result := make(map[string]any)
		for _, key := range rv.MapKeys() {
			result[fmt.Sprintf("%v", key.Interface())] = rv.MapIndex(key).Interface()
		}
		return result
	}
	return nil
}

// MapStrStr 转换为 map[string]string
func (v *Var) MapStrStr() map[string]string {
	m := v.Map()
	if m == nil {
		return nil
	}
	result := make(map[string]string, len(m))
	for k, val := range m {
		result[k] = NewVar(val).String()
	}
	return result
}

// Scan 扫描到目标结构
func (v *Var) Scan(pointer any) error {
	if v.IsNil() {
		return nil
	}
	data, err := json.Marshal(v.value)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, pointer)
}

// MarshalJSON 实现 json.Marshaler
func (v *Var) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Val())
}

// UnmarshalJSON 实现 json.Unmarshaler
func (v *Var) UnmarshalJSON(b []byte) error {
	var i any
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	v.Set(i)
	return nil
}

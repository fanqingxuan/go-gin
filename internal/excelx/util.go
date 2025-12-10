package excelx

import (
	"fmt"
	"reflect"
)

// StructsToRows 将结构体切片转换为 [][]any 数据行
// data: 结构体切片
// fields: 要导出的字段名（按顺序），为空则导出所有字段
func StructsToRows(data any, fields ...string) [][]any {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return nil
	}

	if v.Len() == 0 {
		return nil
	}

	rows := make([][]any, 0, v.Len())

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Struct {
			continue
		}

		var row []any
		if len(fields) > 0 {
			// 按指定字段顺序导出
			for _, fieldName := range fields {
				field := elem.FieldByName(fieldName)
				if field.IsValid() {
					row = append(row, field.Interface())
				} else {
					row = append(row, nil)
				}
			}
		} else {
			// 导出所有字段
			t := elem.Type()
			for j := 0; j < elem.NumField(); j++ {
				if t.Field(j).IsExported() {
					row = append(row, elem.Field(j).Interface())
				}
			}
		}
		rows = append(rows, row)
	}

	return rows
}

// StructsToStringRows 将结构体切片转换为 [][]string 数据行（适用于 CSV）
// data: 结构体切片
// fields: 要导出的字段名（按顺序），为空则导出所有字段
func StructsToStringRows(data any, fields ...string) [][]string {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return nil
	}

	if v.Len() == 0 {
		return nil
	}

	rows := make([][]string, 0, v.Len())

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Struct {
			continue
		}

		var row []string
		if len(fields) > 0 {
			// 按指定字段顺序导出
			for _, fieldName := range fields {
				field := elem.FieldByName(fieldName)
				if field.IsValid() {
					row = append(row, formatValue(field))
				} else {
					row = append(row, "")
				}
			}
		} else {
			// 导出所有字段
			t := elem.Type()
			for j := 0; j < elem.NumField(); j++ {
				if t.Field(j).IsExported() {
					row = append(row, formatValue(elem.Field(j)))
				}
			}
		}
		rows = append(rows, row)
	}

	return rows
}

// formatValue 将 reflect.Value 转换为字符串
func formatValue(v reflect.Value) string {
	if !v.IsValid() {
		return ""
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return ""
		}
		return formatValue(v.Elem())
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

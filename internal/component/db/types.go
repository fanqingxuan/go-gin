package db

import "go-gin/internal/g"

// Value 数据表记录值
type Value = *g.Var

// Record 数据表记录键值对
type Record map[string]Value

// Result 数据表记录列表
type Result []Record

// Map 将 Record 转换为 map[string]any
func (r Record) Map() map[string]any {
	if r == nil {
		return nil
	}
	m := make(map[string]any, len(r))
	for k, v := range r {
		m[k] = v.Val()
	}
	return m
}

// IsEmpty 判断 Record 是否为空
func (r Record) IsEmpty() bool {
	return len(r) == 0
}

// Maps 将 Result 转换为 []map[string]any
func (r Result) Maps() []map[string]any {
	if r == nil {
		return nil
	}
	list := make([]map[string]any, len(r))
	for i, record := range r {
		list[i] = record.Map()
	}
	return list
}

// IsEmpty 判断 Result 是否为空
func (r Result) IsEmpty() bool {
	return len(r) == 0
}

// Len 返回 Result 长度
func (r Result) Len() int {
	return len(r)
}

// toRecord 将 map[string]any 转换为 Record
func toRecord(m map[string]any) Record {
	if m == nil {
		return nil
	}
	r := make(Record, len(m))
	for k, v := range m {
		r[k] = g.NewVar(v)
	}
	return r
}

// toResult 将 []map[string]any 转换为 Result
func toResult(list []map[string]any) Result {
	if list == nil {
		return nil
	}
	r := make(Result, len(list))
	for i, m := range list {
		r[i] = toRecord(m)
	}
	return r
}

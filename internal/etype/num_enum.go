package etype

import (
	"fmt"
	"go-gin/internal/g"
)

type NumEnum int

type INumEnum[T g.IntScalar] interface {
	Equal(a T) bool
	fmt.Stringer
	fmt.Formatter
}

func Equal[T g.IntScalar](a, b T) bool {
	return a == b
}

// Format 实现 fmt.Formatter 接口
func Format[T g.IntScalar](f fmt.State, n T) {
	// 对于所有的默认打印请求（如fmt.Println调用），使用text(xx)格式
	fmt.Fprintf(f, "%d", int(n))
}

func ParseStringFromMap[K g.IntScalar](n K, m map[K]string) string {
	if str, ok := m[n]; ok {
		return str
	}
	return "unknown"
}

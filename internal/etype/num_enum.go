package etype

import (
	"fmt"
	"go-gin/internal/g"
)

type INumEnum interface {
	Equal(a, b any) bool
	fmt.Stringer
	fmt.Formatter
}

type NumEnum int

func Equal[V g.IntScalar](a, b V) bool {
	return a == b
}

// Format 实现 fmt.Formatter 接口
func Format[V g.IntScalar](f fmt.State, n V) {
	// 对于所有的默认打印请求（如fmt.Println调用），使用text(xx)格式
	fmt.Fprintf(f, "%d", int(n))
}

func String[K g.IntScalar](n K, m map[K]string) string {
	if str, ok := m[n]; ok {
		return str
	}
	return "unknown"
}

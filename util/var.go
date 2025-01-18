package util

// Scalar 接口用于限制参数为标量类型或基于标量的自定义类型
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~bool | ~string // 允许基于标量类型的自定义类型
}

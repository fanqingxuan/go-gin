package util

// Scalar 接口用于限制参数为标量类型或基于标量的自定义类型
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~bool | ~string // 允许基于标量类型的自定义类型
}

// IsTrue 检查给定的标量值是否为真
func IsTrue[T Scalar](value T) bool {
	var zeroValue T
	return value != zeroValue // 检查是否为零值
}

// IsFalse 检查给定的标量值是否为假
func IsFalse[T Scalar](value T) bool {
	return !IsTrue(value)
}

// When 执行给定函数,如果给定的标量值为真
func When[T Scalar](value T, f func()) {
	if IsTrue(value) {
		f()
	}
}

// Unless 执行给定函数,如果给定的标量值为假
func Unless[T Scalar](value T, f func()) {
	if IsFalse(value) {
		f()
	}
}

// Value 返回条件表达式的值
func Value[T Scalar, U any](condition T, trueValue U, falseValue U) U {
	if IsTrue(condition) {
		return trueValue
	}
	return falseValue
}

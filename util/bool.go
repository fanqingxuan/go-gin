package util

import "go-gin/internal/g"

// IsTrue 检查给定的标量值是否为真
func IsTrue[T g.Scalar](value T) bool {
	var zeroValue T
	return value != zeroValue // 检查是否为零值
}

// IsFalse 检查给定的标量值是否为假
func IsFalse[T g.Scalar](value T) bool {
	return !IsTrue(value)
}

// WhenFunc 执行给定函数,如果给定的标量值为真
func WhenFunc[T g.Scalar](value T, f func()) {
	if IsTrue(value) {
		f()
	}
}

// When 返回条件表达式的值
func When[T g.Scalar, U any](condition T, trueValue U, falseValue U) U {
	if IsTrue(condition) {
		return trueValue
	}
	return falseValue
}

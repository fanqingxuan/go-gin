package util

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

// Conditional 返回条件表达式的值
func Conditional[T Scalar, U any](condition T, trueValue U, falseValue U) U {
	if IsTrue(condition) {
		return trueValue
	}
	return falseValue
}

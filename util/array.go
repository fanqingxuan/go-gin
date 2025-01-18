package util

// InArray 检查某个元素是否在切片中
func InArray[T Scalar](item T, slice []T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

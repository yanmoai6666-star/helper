package calculator

// 计算两个数之和
func Add(a, b int) int {
	return a + b
}

// 计算两个数之差
func Subtract(a, b int) int {
	return a - b
}

// 计算两个数之积
func Multiply(a, b int) int {
	return a * b
}

// 计算两个数之商
func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

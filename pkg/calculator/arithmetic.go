package calculator

// 计算两个整数的平方和
func SquareSum(a, b int) int {
	return a*a + b*b
}

// 计算两个整数的平方差
func SquareDifference(a, b int) int {
	return a*a - b*b
}

// 计算一个数的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 计算一个数的三次方
func Cube(x int) int {
	return x * x * x
}

// 计算两个数的最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算两个数的最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

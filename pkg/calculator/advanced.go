package calculator

import "math"

// 计算阶乘
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 计算斐波那契数列第n项
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 计算平方根（整数近似）
func SquareRoot(x int) int {
	return int(math.Sqrt(float64(x)))
}

// 计算对数（以e为底）
func NaturalLog(x float64) float64 {
	return math.Log(x)
}

// 计算幂运算
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

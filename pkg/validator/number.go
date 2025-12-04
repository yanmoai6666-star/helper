package validator

// 检查整数是否在指定范围内
func IsIntBetween(n, min, max int) ValidationResult {
	if n >= min && n <= max {
		return ValidResult()
	}
	return InvalidResult("整数必须在指定范围内")
}

// 检查整数是否大于等于指定值
func IsIntMin(n, min int) ValidationResult {
	if n >= min {
		return ValidResult()
	}
	return InvalidResult("整数必须大于等于指定值")
}

// 检查整数是否小于等于指定值
func IsIntMax(n, max int) ValidationResult {
	if n <= max {
		return ValidResult()
	}
	return InvalidResult("整数必须小于等于指定值")
}

// 检查浮点数是否在指定范围内
func IsFloatBetween(f, min, max float64) ValidationResult {
	if f >= min && f <= max {
		return ValidResult()
	}
	return InvalidResult("浮点数必须在指定范围内")
}

// 检查浮点数是否大于等于指定值
func IsFloatMin(f, min float64) ValidationResult {
	if f >= min {
		return ValidResult()
	}
	return InvalidResult("浮点数必须大于等于指定值")
}

// 检查浮点数是否小于等于指定值
func IsFloatMax(f, max float64) ValidationResult {
	if f <= max {
		return ValidResult()
	}
	return InvalidResult("浮点数必须小于等于指定值")
}

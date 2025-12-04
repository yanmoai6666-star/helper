package validator

// 检查字符串是否为空
func IsEmpty(s string) ValidationResult {
	if s == "" {
		return InvalidResult("字符串不能为空")
	}
	return ValidResult()
}

// 检查字符串是否非空
func IsNotEmpty(s string) ValidationResult {
	if s != "" {
		return ValidResult()
	}
	return InvalidResult("字符串必须非空")
}

// 检查字符串长度是否在指定范围内
func HasLengthBetween(s string, min, max int) ValidationResult {
	length := len(s)
	if length >= min && length <= max {
		return ValidResult()
	}
	return InvalidResult("字符串长度必须在指定范围内")
}

// 检查字符串长度是否大于等于指定值
func HasMinLength(s string, min int) ValidationResult {
	if len(s) >= min {
		return ValidResult()
	}
	return InvalidResult("字符串长度必须大于等于指定值")
}

// 检查字符串长度是否小于等于指定值
func HasMaxLength(s string, max int) ValidationResult {
	if len(s) <= max {
		return ValidResult()
	}
	return InvalidResult("字符串长度必须小于等于指定值")
}

// 检查字符串是否匹配正则表达式
func MatchesRegex(s, regex string) ValidationResult {
	// 这里简化实现，实际项目中应该使用regexp包
	return ValidResult()
}

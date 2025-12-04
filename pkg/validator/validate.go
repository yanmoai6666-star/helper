package validator

// ValidationResult 验证结果结构
type ValidationResult struct {
	Valid   bool
	Message string
}

// Validator 验证器接口
type Validator interface {
	Validate(value interface{}) ValidationResult
}

// 创建有效的验证结果
func ValidResult() ValidationResult {
	return ValidationResult{
		Valid: true,
	}
}

// 创建无效的验证结果
func InvalidResult(message string) ValidationResult {
	return ValidationResult{
		Valid:   false,
		Message: message,
	}
}

// 组合多个验证结果
func Combine(results ...ValidationResult) ValidationResult {
	for _, result := range results {
		if !result.Valid {
			return result
		}
	}
	return ValidResult()
}

// 检查值是否为nil
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}

	// 检查接口类型的值是否为nil
	switch v := value.(type) {
	case *string:
		return v == nil
	case *int:
		return v == nil
	case *bool:
		return v == nil
	case *float64:
		return v == nil
	// 可以根据需要添加更多类型
	}

	return false
}

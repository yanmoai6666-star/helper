package validator

// StructValidator 结构体验证器接口
type StructValidator interface {
	ValidateStruct(s interface{}) ValidationResult
}

// FieldValidator 字段验证器接口
type FieldValidator interface {
	ValidateField(fieldName string, value interface{}) ValidationResult
}

// ValidationRule 验证规则结构
type ValidationRule struct {
	Field    string
	Validator Validator
	Message  string
}

// 创建字段验证规则
func NewRule(field string, validator Validator, message string) ValidationRule {
	return ValidationRule{
		Field:    field,
		Validator: validator,
		Message:  message,
	}
}

// RuleSet 验证规则集合
type RuleSet []ValidationRule

// 添加验证规则
func (rs *RuleSet) Add(rule ValidationRule) {
	*rs = append(*rs, rule)
}

// 应用验证规则集合
func (rs RuleSet) Apply(value interface{}) ValidationResult {
	for _, rule := range rs {
		result := rule.Validator.Validate(value)
		if !result.Valid {
			return InvalidResult(rule.Message)
		}
	}
	return ValidResult()
}

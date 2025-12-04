package stringutil

import (
	"regexp"
	"strings"
)

// 检查字符串是否为有效的电子邮件格式
func IsEmail(s string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(s)
}

// 检查字符串是否为有效的URL格式
func IsURL(s string) bool {
	urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(s)
}

// 检查字符串是否只包含字母
func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

// 检查字符串是否只包含字母和数字
func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// 检查字符串长度是否在指定范围内
func HasLengthBetween(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}

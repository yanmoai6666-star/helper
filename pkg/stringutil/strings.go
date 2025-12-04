package stringutil

// 检查字符串是否为空
func IsEmpty(s string) bool {
	return s == ""
}

// IsNotEmpty 检查字符串是否非空
func IsNotEmpty(s string) bool {
	return s != ""
}

// 检查字符串是否为空白
func IsBlank(s string) bool {
	for i := range s {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && s[i] != '\r' {
			return false
		}
	}
	return true
}

// IsNotBlank 检查字符串是否非空白
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// 检查两个字符串是否相等（忽略大小写）
func EqualsIgnoreCase(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		ac := a[i]
		bc := b[i]
		if ac >= 'A' && ac <= 'Z' {
			ac += 'a' - 'A'
		}
		if bc >= 'A' && bc <= 'Z' {
			bc += 'a' - 'A'
		}
		if ac != bc {
			return false
		}
	}
	return true
}

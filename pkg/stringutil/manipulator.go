package stringutil

import (
	"strings"
)

// 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 截取字符串，从指定位置开始
func Substring(s string, start int) string {
	if start < 0 {
		start = 0
	}
	if start >= len(s) {
		return ""
	}
	return s[start:]
}

// 截取字符串，从指定位置开始到指定位置结束
func SubstringBetween(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return ""
	}
	return s[start:end]
}

// 替换字符串中的子串
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// 分割字符串
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// 连接字符串切片
func Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// 计算字符串中特定子串出现的次数
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

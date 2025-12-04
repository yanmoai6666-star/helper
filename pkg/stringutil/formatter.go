package stringutil

import (
	"fmt"
	"strings"
)

// 格式化字符串，将第一个字符转为大写
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// 格式化字符串，将每个单词的第一个字符转为大写
func Title(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = Capitalize(word)
	}
	return strings.Join(words, " ")
}

// 移除字符串两端的空白字符
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// 移除字符串左侧的空白字符
func TrimLeft(s string) string {
	return strings.TrimLeftFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
}

// 移除字符串右侧的空白字符
func TrimRight(s string) string {
	return strings.TrimRightFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
}

// 重复字符串指定次数
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

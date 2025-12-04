package timeutil

import (
	"time"
)

// 标准时间格式常量
const (
	FormatDate     = "2006-01-02"
	FormatTime     = "15:04:05"
	FormatDateTime = "2006-01-02 15:04:05"
	FormatISO8601  = "2006-01-02T15:04:05Z07:00"
)

// 将时间格式化为日期字符串
func FormatDate(t time.Time) string {
	return t.Format(FormatDate)
}

// 将时间格式化为时间字符串
func FormatTime(t time.Time) string {
	return t.Format(FormatTime)
}

// 将时间格式化为日期时间字符串
func FormatDateTime(t time.Time) string {
	return t.Format(FormatDateTime)
}

// 将时间格式化为ISO8601格式
func FormatISO8601(t time.Time) string {
	return t.Format(FormatISO8601)
}

// 使用自定义格式格式化时间
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

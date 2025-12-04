package timeutil

import (
	"time"
)

// 创建秒级持续时间
func Seconds(seconds int64) time.Duration {
	return time.Duration(seconds) * time.Second
}

// 创建分钟级持续时间
func Minutes(minutes int64) time.Duration {
	return time.Duration(minutes) * time.Minute
}

// 创建小时级持续时间
func Hours(hours int64) time.Duration {
	return time.Duration(hours) * time.Hour
}

// 创建天级持续时间
func Days(days int64) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}

// 获取持续时间的总秒数
func TotalSeconds(d time.Duration) float64 {
	return d.Seconds()
}

// 获取持续时间的总分钟数
func TotalMinutes(d time.Duration) float64 {
	return d.Minutes()
}

// 获取持续时间的总小时数
func TotalHours(d time.Duration) float64 {
	return d.Hours()
}

// 获取持续时间的总天数
func TotalDays(d time.Duration) float64 {
	return d.Hours() / 24
}

// 格式化持续时间为人类可读的字符串
func FormatDuration(d time.Duration) string {
	if d < time.Second {
		return d.Round(time.Millisecond).String()
	}
	return d.Round(time.Second).String()
}

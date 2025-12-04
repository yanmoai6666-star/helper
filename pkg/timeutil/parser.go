package timeutil

import (
	"time"
)

// 解析日期字符串为时间对象
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(FormatDate, dateStr)
}

// 解析时间字符串为时间对象
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(FormatTime, timeStr)
}

// 解析日期时间字符串为时间对象
func ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.Parse(FormatDateTime, dateTimeStr)
}

// 解析ISO8601格式的时间字符串
func ParseISO8601(isoStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, isoStr)
}

// 使用自定义格式解析时间字符串
func Parse(timeStr, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

// 解析时间戳（秒）为时间对象
func ParseTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// 解析时间戳（毫秒）为时间对象
func ParseTimestampMillis(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

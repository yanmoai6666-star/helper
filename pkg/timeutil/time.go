package timeutil

import (
	"time"
)

// 获取当前时间
func Now() time.Time {
	return time.Now()
}

// 获取当前时间的UTC表示
func NowUTC() time.Time {
	return time.Now().UTC()
}

// 获取当前时间戳（秒）
func Timestamp() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳（毫秒）
func TimestampMillis() int64 {
	return time.Now().UnixMilli()
}

// 获取当前时间戳（微秒）
func TimestampMicros() int64 {
	return time.Now().UnixMicro()
}

// 获取当前时间戳（纳秒）
func TimestampNanos() int64 {
	return time.Now().UnixNano()
}

// 将时间戳转换为时间对象
func FromTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

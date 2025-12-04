package logger

import (
	"fmt"
	"time"
)

// 日志记录器结构

type Logger struct {
	Level string
}

// 初始化日志记录器
func NewLogger(level string) *Logger {
	return &Logger{
		Level: level,
	}
}

// 记录信息级别日志
func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

// 记录警告级别日志
func (l *Logger) Warn(message string) {
	l.log("WARN", message)
}

// 记录错误级别日志
func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

// 实际日志输出函数
func (l *Logger) log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
}

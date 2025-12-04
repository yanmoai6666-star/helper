package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// 日志格式化器接口
type Formatter interface {
	Format(level LogLevel, message string, fields map[string]interface{}) string
}

// 文本格式化器
type TextFormatter struct {
	TimestampFormat string
}

// JSON格式化器
type JsonFormatter struct {
	Indent bool
}

// 创建文本格式化器
func NewTextFormatter(timestampFormat string) *TextFormatter {
	if timestampFormat == "" {
		timestampFormat = "2006-01-02 15:04:05"
	}
	return &TextFormatter{
		TimestampFormat: timestampFormat,
	}
}

// 创建JSON格式化器
func NewJsonFormatter(indent bool) *JsonFormatter {
	return &JsonFormatter{
		Indent: indent,
	}
}

// 文本格式化实现
func (f *TextFormatter) Format(level LogLevel, message string, fields map[string]interface{}) string {
	timestamp := time.Now().Format(f.TimestampFormat)
	logLine := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
	
	for k, v := range fields {
		logLine += fmt.Sprintf(" %s=%v", k, v)
	}
	
	return logLine
}

// JSON格式化实现
func (f *JsonFormatter) Format(level LogLevel, message string, fields map[string]interface{}) string {
	logEntry := map[string]interface{}{
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"level":     level,
		"message":   message,
	}
	
	for k, v := range fields {
		logEntry[k] = v
	}
	
	var data []byte
	var err error
	
	if f.Indent {
		data, err = json.MarshalIndent(logEntry, "", "  ")
	} else {
		data, err = json.Marshal(logEntry)
	}
	
	if err != nil {
		return fmt.Sprintf("[{"error": "%s"}]", err.Error())
	}
	
	return string(data)
}

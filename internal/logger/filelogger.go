package logger

import (
	"os"
	"sync"
	"time"
)

// 文件日志记录器
type FileLogger struct {
	Logger
	file     *os.File
	mutex    sync.Mutex
	config   *LoggerConfig
	formatter Formatter
}

// 创建文件日志记录器
func NewFileLogger(config *LoggerConfig) (*FileLogger, error) {
	file, err := os.OpenFile(
		config.Filepath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, err
	}

	fl := &FileLogger{
		Logger: Logger{
			Level: string(config.Level),
		},
		file:     file,
		config:   config,
		formatter: NewTextFormatter("2006-01-02 15:04:05"),
	}

	return fl, nil
}

// 关闭文件日志记录器
func (fl *FileLogger) Close() error {
	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	if fl.file != nil {
		return fl.file.Close()
	}
	return nil
}

// 记录信息级别日志到文件
func (fl *FileLogger) Info(message string) {
	fl.logToFile(LevelInfo, message)
}

// 记录警告级别日志到文件
func (fl *FileLogger) Warn(message string) {
	fl.logToFile(LevelWarn, message)
}

// 记录错误级别日志到文件
func (fl *FileLogger) Error(message string) {
	fl.logToFile(LevelError, message)
}

// 记录日志到文件
func (fl *FileLogger) logToFile(level LogLevel, message string) {
	if !fl.shouldLog(level) {
		return
	}

	fl.mutex.Lock()
	defer fl.mutex.Unlock()

	formatted := fl.formatter.Format(level, message, nil)
	_, err := fl.file.WriteString(formatted + "\n")
	if err != nil {
		// 错误处理：可以记录到标准错误
		os.Stderr.WriteString(fmt.Sprintf("Failed to write log: %s\n", err.Error()))
	}
}

// 检查是否应该记录该级别日志
func (fl *FileLogger) shouldLog(level LogLevel) bool {
	levels := map[LogLevel]int{
		LevelDebug: 0,
		LevelInfo:  1,
		LevelWarn:  2,
		LevelError: 3,
		LevelFatal: 4,
	}

	current, exists := levels[level]
	if !exists {
		return false
	}

	loggerLevel, exists := levels[LogLevel(fl.Level)]
	if !exists {
		return true
	}

	return current >= loggerLevel
}

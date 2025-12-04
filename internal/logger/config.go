package logger

// 日志级别类型
type LogLevel string

// 日志级别常量
const (
	LevelDebug LogLevel = "DEBUG"
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
	LevelFatal LogLevel = "FATAL"
)

// 日志配置结构
type LoggerConfig struct {
	Level     LogLevel
	Format    string
	Output    string
	Filepath  string
	MaxSize   int64
	MaxAge    int
	MaxBackup int
}

// 创建默认日志配置
func DefaultConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:    LevelInfo,
		Format:   "text",
		Output:   "stdout",
		MaxSize:  1024 * 1024 * 100, // 100MB
		MaxAge:   7,                // 7天
		MaxBackup: 5,               // 5个备份
	}
}

// 设置日志级别
func (c *LoggerConfig) WithLevel(level LogLevel) *LoggerConfig {
	c.Level = level
	return c
}

// 设置日志格式
func (c *LoggerConfig) WithFormat(format string) *LoggerConfig {
	c.Format = format
	return c
}

// 设置文件输出路径
func (c *LoggerConfig) WithFile(filepath string) *LoggerConfig {
	c.Output = "file"
	c.Filepath = filepath
	return c
}

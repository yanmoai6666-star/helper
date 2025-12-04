package config

// Config 配置结构
type Config struct {
	AppName     string
	Environment string
	Port        int
	Debug       bool
	Database    DatabaseConfig
	Redis       RedisConfig
}

// DatabaseConfig 数据库配置结构
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// RedisConfig Redis配置结构
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// 创建默认配置
func DefaultConfig() *Config {
	return &Config{
		AppName:     "helper",
		Environment: "development",
		Port:        8080,
		Debug:       true,
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "",
			Name:     "helper_db",
		},
		Redis: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
}

// GetEnvironment 获取当前环境
func (c *Config) GetEnvironment() string {
	return c.Environment
}

// IsDevelopment 检查是否为开发环境
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction 检查是否为生产环境
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

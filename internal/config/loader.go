package config

import (
	"os"
	"strconv"
)

// ConfigLoader 配置加载器接口
type ConfigLoader interface {
	Load() (*Config, error)
}

// EnvironmentLoader 从环境变量加载配置
type EnvironmentLoader struct{}

// 创建环境变量加载器
func NewEnvironmentLoader() *EnvironmentLoader {
	return &EnvironmentLoader{}
}

// Load 从环境变量加载配置
func (l *EnvironmentLoader) Load() (*Config, error) {
	config := DefaultConfig()

	// 加载应用配置
	if appName := os.Getenv("APP_NAME"); appName != "" {
		config.AppName = appName
	}

	if env := os.Getenv("ENVIRONMENT"); env != "" {
		config.Environment = env
	}

	if portStr := os.Getenv("PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			config.Port = port
		}
	}

	if debugStr := os.Getenv("DEBUG"); debugStr != "" {
		debug, err := strconv.ParseBool(debugStr)
		if err == nil {
			config.Debug = debug
		}
	}

	// 加载数据库配置
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}

	if dbPortStr := os.Getenv("DB_PORT"); dbPortStr != "" {
		dbPort, err := strconv.Atoi(dbPortStr)
		if err == nil {
			config.Database.Port = dbPort
		}
	}

	if dbUser := os.Getenv("DB_USERNAME"); dbUser != "" {
		config.Database.Username = dbUser
	}

	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		config.Database.Password = dbPass
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.Database.Name = dbName
	}

	// 加载Redis配置
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		config.Redis.Host = redisHost
	}

	if redisPortStr := os.Getenv("REDIS_PORT"); redisPortStr != "" {
		redisPort, err := strconv.Atoi(redisPortStr)
		if err == nil {
			config.Redis.Port = redisPort
		}
	}

	if redisPass := os.Getenv("REDIS_PASSWORD"); redisPass != "" {
		config.Redis.Password = redisPass
	}

	if redisDBStr := os.Getenv("REDIS_DB"); redisDBStr != "" {
		redisDB, err := strconv.Atoi(redisDBStr)
		if err == nil {
			config.Redis.DB = redisDB
		}
	}

	return config, nil
}

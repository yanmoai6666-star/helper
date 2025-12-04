package config

import (
	"fmt"
	"strings"
)

// ConfigValidator 配置验证器接口
type ConfigValidator interface {
	Validate(config *Config) error
}

// DefaultConfigValidator 默认配置验证器
type DefaultConfigValidator struct{}

// 创建默认配置验证器
func NewDefaultConfigValidator() *DefaultConfigValidator {
	return &DefaultConfigValidator{}
}

// Validate 验证配置
func (v *DefaultConfigValidator) Validate(config *Config) error {
	var errors []string

	// 验证应用配置
	if config.AppName == "" {
		errors = append(errors, "应用名称不能为空")
	}

	if config.Environment == "" {
		errors = append(errors, "环境不能为空")
	} else if config.Environment != "development" && config.Environment != "staging" && config.Environment != "production" {
		errors = append(errors, "环境必须是 development、staging 或 production")
	}

	if config.Port <= 0 || config.Port > 65535 {
		errors = append(errors, "端口必须在 1-65535 之间")
	}

	// 验证数据库配置
	if config.Database.Host == "" {
		errors = append(errors, "数据库主机不能为空")
	}

	if config.Database.Port <= 0 || config.Database.Port > 65535 {
		errors = append(errors, "数据库端口必须在 1-65535 之间")
	}

	if config.Database.Username == "" {
		errors = append(errors, "数据库用户名不能为空")
	}

	if config.Database.Name == "" {
		errors = append(errors, "数据库名称不能为空")
	}

	// 验证Redis配置
	if config.Redis.Host == "" {
		errors = append(errors, "Redis主机不能为空")
	}

	if config.Redis.Port <= 0 || config.Redis.Port > 65535 {
		errors = append(errors, "Redis端口必须在 1-65535 之间")
	}

	if len(errors) > 0 {
		return fmt.Errorf("配置验证失败: %s", strings.Join(errors, ", "))
	}

	return nil
}

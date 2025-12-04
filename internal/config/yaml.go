package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// YAMLLoader YAML配置加载器
type YAMLLoader struct {
	Path string
}

// 创建YAML配置加载器
func NewYAMLLoader(path string) *YAMLLoader {
	return &YAMLLoader{Path: path}
}

// Load 从YAML文件加载配置
func (l *YAMLLoader) Load() (*Config, error) {
	config := DefaultConfig()

	file, err := ioutil.ReadFile(l.Path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

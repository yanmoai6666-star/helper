package config

// ConfigProvider 配置提供者接口
type ConfigProvider interface {
	GetConfig() (*Config, error)
	Reload() error
}

// StaticConfigProvider 静态配置提供者
type StaticConfigProvider struct {
	config *Config
}

// 创建静态配置提供者
func NewStaticConfigProvider(config *Config) *StaticConfigProvider {
	return &StaticConfigProvider{config: config}
}

// GetConfig 获取配置
func (p *StaticConfigProvider) GetConfig() (*Config, error) {
	return p.config, nil
}

// Reload 重新加载配置
func (p *StaticConfigProvider) Reload() error {
	return nil
}

// DynamicConfigProvider 动态配置提供者
type DynamicConfigProvider struct {
	loader    ConfigLoader
	validator ConfigValidator
	config    *Config
}

// 创建动态配置提供者
func NewDynamicConfigProvider(loader ConfigLoader, validator ConfigValidator) *DynamicConfigProvider {
	return &DynamicConfigProvider{
		loader:    loader,
		validator: validator,
	}
}

// GetConfig 获取配置
func (p *DynamicConfigProvider) GetConfig() (*Config, error) {
	if p.config == nil {
		err := p.Reload()
		if err != nil {
			return nil, err
		}
	}
	return p.config, nil
}

// Reload 重新加载配置
func (p *DynamicConfigProvider) Reload() error {
	config, err := p.loader.Load()
	if err != nil {
		return err
	}

	err = p.validator.Validate(config)
	if err != nil {
		return err
	}

	p.config = config
	return nil
}

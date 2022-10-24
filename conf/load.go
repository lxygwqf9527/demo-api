package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 从配置文件中加载
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path: %s, %s", filePath, err)
	}
	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}

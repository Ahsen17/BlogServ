package config

import (
	"BlogServ/logger"
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Address string `toml:"address"`
		Port    int    `toml:"port"`
		Timeout int    `toml:"timeout"`
	}

	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Database string `toml:"database"`
	}

	Common struct {
		Environment string `toml:"environment"`
	}

	Logging struct {
		Level   string `toml:"level"`
		Enabled bool   `toml:"enabled"`
	}
}

func NewConfig() *Config {
	data, err := ioutil.ReadFile("../config.toml")
	if err != nil {
		logger.Fatal("打开配置文件global.toml失败,请检查文件是否存在")
	}

	var config Config
	if err = toml.Unmarshal(data, &config); err != nil {
		logger.Fatal("解析配置失败,请检查配置是否正确")
	}

	return &config
}

package config

import (
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
	"log"
	"sync"
)

type Server struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
	Timeout int    `toml:"timeout"`
}

type Database struct {
	Driver   string `toml:"driver"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	MaxConn  int    `toml:"maxConn"`
	IdleConn int    `toml:"idleConn"`
}

type Common struct {
	Environment string `toml:"environment"`
}

type Logging struct {
	Level   string `toml:"level"`
	Enabled bool   `toml:"enabled"`
}

type Config struct {
	Server   Server
	Database Database
	Common   Common
	Logging  Logging
}

var config *Config
var once sync.Once

var BaseDir = "/etc/blogserv"

// FetchConfig 单例配置对象
func FetchConfig() *Config {
	once.Do(func() {
		data, err := ioutil.ReadFile(BaseDir + "/config.toml")
		if err != nil {
			log.Fatal("打开配置文件global.toml失败,请检查文件是否存在")
		}

		if err = toml.Unmarshal(data, &config); err != nil {
			log.Fatal("解析配置失败,请检查配置是否正确")
		}

		log.Println("初始化系统配置成功")
	})

	return config
}

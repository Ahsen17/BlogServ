package config

import (
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
	"log"
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

type Cache struct {
	Host     string `toml:"host"`
	Port     uint   `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
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
	Cache    Cache
	Common   Common
	Logging  Logging
}

const (
	DEVENV  = "dev"
	TESTENV = "test"
	PRODENV = "prod"
)

var (
	config *Config
)

//global variables
var (
	projectEnv = DEVENV
	configFile string

	ConfDir = "/etc/blogserv"
)

// NewConfig 单例配置对象
func init() {
	InitGlobalVar()

	// 本地开发使用config.toml配置,拉取源码时需要修改为global.toml
	data, err := ioutil.ReadFile(ConfDir + configFile)
	if err != nil {
		log.Fatal("打开配置文件失败,请检查文件是否存在")
	}

	if err = toml.Unmarshal(data, &config); err != nil {
		log.Fatal("解析配置失败,请检查配置是否正确")
	}

	log.Println("初始化系统配置成功")
}

func InitGlobalVar() {
	if projectEnv == DEVENV {
		// 开发环境
		configFile = "/config.toml"
	} else if projectEnv == TESTENV {
		// 测试环境
		configFile = ""
	} else if projectEnv == PRODENV {
		// 生产环境
		configFile = "/global.toml"
	} else {
		configFile = ""
	}
}

func AllConfig() *Config {
	return config
}

func ServerConfig() *Server {
	return &config.Server
}

// DBConfig 获取数据库配置
func DBConfig() *Database {
	return &config.Database
}

func CacheConfig() *Cache {
	return &config.Cache
}

// CommonConfig 获取Common配置
func CommonConfig() *Common {
	return &config.Common
}

// LoggingConfig 获取日志配置
func LoggingConfig() *Logging {
	return &config.Logging
}

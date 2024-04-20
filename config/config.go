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
	Level        string `toml:"level"`
	LogFile      string `toml:"logFile"`
	EnableOutput bool   `toml:"enableOutput"`
}

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	Common   Common
	Logging  Logging
}

const (
	DEV  = "dev"
	TEST = "test"
	PROD = "prod"
)

var (
	config *Config
)

// global variables
var (
	ConfFile   string
	ConfDir    string
	ProjectEnv string
)

// NewConfig 单例配置对象
func init() {
	InitGlobalVar()

	// 本地开发使用config.toml配置,拉取源码时需要修改为global.toml
	data, err := ioutil.ReadFile(ConfDir + ConfFile)
	if err != nil {
		log.Fatalf("读取配置文件失败,请检查文件[%s]是否存在", ConfFile)
	}

	if err = toml.Unmarshal(data, &config); err != nil {
		log.Fatal("解析配置失败,请检查配置是否正确")
	}

	log.Println("初始化系统配置成功")
}

func InitGlobalVar() {
	ConfDir = "./"

	ProjectEnv = DEV
	switch ProjectEnv {
	case DEV:
		ConfFile = "dev.toml"
	case TEST:
		ConfFile = "test.toml"
	default:
		ConfFile = "config.toml"
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

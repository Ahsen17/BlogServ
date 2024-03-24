package config

import (
	"fmt"
	"testing"
)

// 测试获取Server相关的配置属性
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	sc := c.Server
	fmt.Println(sc)
}

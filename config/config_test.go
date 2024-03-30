package config

import (
	"fmt"
	"testing"
)

// 测试获取Server相关的配置属性
func TestFetchConfig(t *testing.T) {
	c := NewConfig()
	c2 := NewConfig()
	//sc := c.Server
	fmt.Println(c == c2)
}

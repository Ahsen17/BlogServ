package config

import (
	"fmt"
	"testing"
)

// 测试获取Server相关的配置属性
func TestFetchConfig(t *testing.T) {
	c := AllConfig()
	//sc := c.Server
	fmt.Println(c)
}

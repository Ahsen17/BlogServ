/*
  Package tools
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/6/2 13:25
  @Description: ...
*/

package tools

import (
	"math/rand"
	"time"
)

type RandomTool struct {
}

// RandomString 生成随机字符串
func (rt *RandomTool) RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomNumber 生成随机数字
func (rt *RandomTool) RandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

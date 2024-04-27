/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 用户管理
*/

package models

type User struct {
	Email   string
	Sex     uint
	Address string
}

type UserData struct {
}

type UserMgr struct {
	User     *User
	UserData *UserData
}

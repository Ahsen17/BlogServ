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

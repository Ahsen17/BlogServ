package models

const (
	ANONYMOUS = "anonymous"
	AUTHOR    = "author"
	MANAGER   = "manager"
)

type User struct {
	Email   string
	Sex     uint
	Address string
}

type UserMgr struct {
	User *User
}

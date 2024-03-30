package models

import "gorm.io/gorm"

const (
	ANONYMOUS = "anonymous"
	AUTHOR    = "author"
	MANAGER   = "manager"
)

const (
	ACTIVE     = 1
	DEACTIVATE = 2
	REVOKE     = 3
	BANED      = 9
)

type Account struct {
	gorm.Model
	Username    string
	Password    string
	Token       string
	EnableToken bool
	Status      uint
	User        User

	Role   Role
	RoleID uint
}

type Role struct {
	ID    uint `gorm:"primary_key;"`
	Name  string
	Level string
}

type User struct {
	Email   string
	Sex     uint
	Address string
}

type Author struct {
	gorm.Model
	Account   Account
	AccountID uint
}

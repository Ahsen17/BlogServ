package models

import "gorm.io/gorm"

const (
	ANONYMOUS = "anonymous"
	AUTHOR    = "author"
	MANAGER   = "manager"
)

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

package models

type Role struct {
	ID    uint
	Name  string
	Level string
}

type RoleMgr struct {
	Role *Role
}

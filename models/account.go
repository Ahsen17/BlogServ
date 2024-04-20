package models

import (
	"github.com/ahsen17/BlogServ/src/data"
	"time"
)

const (
	ACTIVE     = 1
	DEACTIVATE = 2
	REVOKE     = 3
	BANED      = 9

	DefaultAccount = "system"
)

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	CreateAt int64  `json:"create_at"`
	CreateBy string `json:"create_by"`
	UpdateAt int64  `json:"update_at"`
	UpdateBy string `json:"update_by"`
}

// Register 账户注册
func (ac *Account) Register() error {
	now := time.Now().Unix()
	regSql := "INSERT INTO `account`(username, password, create_at, create_by, update_at, update_by) values (?,?,?,?,?,?)"
	client := data.DBClient()
	// 虽然用了gorm框架,但最终还是决定用原生sql
	// 需要确定后续是否更换orm框架
	client.Raw(
		regSql,
		ac.Username, ac.Password,
		now, DefaultAccount,
		now, DefaultAccount,
	).Scan(&ac)

	return nil
}

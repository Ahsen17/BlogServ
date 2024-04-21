package models

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

type AccountData struct {
	LastLoginTime int64
	LastLoginIp   string
}

type AccountMgr struct {
	Account     *Account
	AccountData *AccountData

	UserMgr      *UserMgr
	RoleMgr      *RoleMgr
	CustomLogMgr *CustomLogMgr // 用户操作日志
}

// Register 用户注册
func (mgr *AccountMgr) Register() error {
	return nil
}

// Login 用户登录
func (mgr *AccountMgr) Login() error {
	return nil
}

// Logout 用户登出
func (mgr *AccountMgr) Logout() error {
	return nil
}

// Edit 更新用户信息
func (mgr *AccountMgr) Edit() error {
	return nil
}

// Revoke 注销用户
func (mgr *AccountMgr) Revoke() error {
	return nil
}

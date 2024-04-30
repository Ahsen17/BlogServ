/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 角色及权限管理
*/

package models

const (
	TableRole = "role"

	SYSTEM = iota + 1
	ADMIN
	MANAGER
	AUTHOR
	ANONYMOUS
)

var RoleMap = map[int]string{
	SYSTEM:    "系统",
	ADMIN:     "超级用户",
	MANAGER:   "管理员",
	AUTHOR:    "创作者",
	ANONYMOUS: "匿名用户",
}

var RoleLevel = map[int]int{
	SYSTEM:    99,
	ADMIN:     70,
	MANAGER:   30,
	AUTHOR:    10,
	ANONYMOUS: 0,
}

type Role struct {
	Sign  int    `json:"sign"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type RoleMgr struct {
	Role *Role
}

func (role Role) TableName() string {
	return TableRole
}

func (mgr *RoleMgr) GetRolesArray() []string {
	var values []string
	for _, v := range RoleMap {
		values = append(values, v)
	}
	return values
}

func (mgr *RoleMgr) GetRolesMap() map[int]string {
	return RoleMap
}

func (mgr *RoleMgr) Info() {
	role := mgr.Role.Sign
	if role <= 0 {
		return
	}
	mgr.Role.Name = RoleMap[role]
	mgr.Role.Level = RoleLevel[role]
}

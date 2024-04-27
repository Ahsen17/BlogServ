/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 角色及权限管理
*/

package models

const (
	SYSTEM    = "system"
	ADMIN     = "admin"
	MANAGER   = "manager"
	AUTHOR    = "author"
	ANONYMOUS = "anonymous"
)

var RoleMap = map[string]string{
	SYSTEM:    "系统",
	ADMIN:     "超级用户",
	MANAGER:   "管理员",
	AUTHOR:    "创作者",
	ANONYMOUS: "匿名用户",
}

type Role struct {
	ID    uint
	Name  string
	Level int
}

type RoleMgr struct {
	Role *Role
}

func (mgr *RoleMgr) GetRolesArray() []string {
	var values []string
	for _, v := range RoleMap {
		values = append(values, v)
	}
	return values
}

func (mgr *RoleMgr) GetRolesMap() map[string]string {
	return RoleMap
}

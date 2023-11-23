package model

import "go-admin-example/authorize/boot"

func init() {
	boot.Register(&SysRole{}, &SysUserRole{}, &SysRoleMenu{})
}

type SysRole struct {
	Id        int32  `gorm:"type:int AUTO_INCREMENT;primary_key;not null"`
	RoleName  string `gorm:"type:varchar(128);not null;comment:角色名称"`
	Desc      string `gorm:"type:varchar(256);not null;default:'';comment:角色描述"`
	RoleKey   string `gorm:"type:varchar(128);not null;comment:角色标识;index"`
	Status    int32  `gorm:"type:int;not null;default:1;comment:状态 1-正常 2-停用"`
	IsDelete  bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否删除 true-是 false-否"`
	IsDefault bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否默认 true-是 false-否"`
	CreatedAt int64  `gorm:"type:bigint;not null;default:0;comment:创建时间"`
	CreatedBy string `gorm:"type:varchar(128);not null;default:'';comment:创建人"`
	UpdatedAt int64  `gorm:"type:bigint;not null;default:0;comment:更新时间"`
	UpdatedBy string `gorm:"type:varchar(128);not null;default:'';comment:更新人"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

func (SysRole) TableComment() string {
	return "系统角色表"
}

type SysRoleMenu struct {
	Id     int32 `gorm:"type:int AUTO_INCREMENT;primary_key;not null"`
	RoleId int32 `gorm:"type:int;not null;comment:角色id;index"`
	MenuId int32 `gorm:"type:int;not null;comment:菜单id;index"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (SysRoleMenu) TableComment() string {
	return "系统角色菜单关联表"
}

type SysUserRole struct {
	Id     int32 `gorm:"type:int AUTO_INCREMENT;primary_key;not null"`
	UserId int32 `gorm:"type:int;not null;comment:用户id;index"`
	RoleId int32 `gorm:"type:int;not null;comment:角色id;index"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}

func (SysUserRole) TableComment() string {
	return "系统用户角色关联表"
}

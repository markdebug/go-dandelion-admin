package model

import "go-admin-example/authorize/boot"

func init() {
	boot.Register(&SysMenu{}, &SysApi{}, &SysMenuApi{})
}

type SysMenu struct {
	Id        int32  `gorm:"type:int Auto_Increment;primary_key;not null"`
	Code      string `gorm:"type:varchar(128);not null;default:'';comment:菜单编号"`
	Title     string `gorm:"type:varchar(128);not null;default:'';comment:显示名称"`
	Icon      string `gorm:"type:varchar(128);not null;default:'';comment:菜单图标"`
	MenuType  string `gorm:"type:varchar(64);not null;default:'';comment:菜单类型"` // module 模块 menu 菜单 3 按钮
	ParentId  int32  `gorm:"type:int;not null;default:0;comment:父菜单ID"`
	Component string `gorm:"type:varchar(256);not null;default:'';comment:前端组件"`
	Path      string `gorm:"type:varchar(256);not null;default:'';comment:路由地址"`
	Sort      int32  `gorm:"type:int;not null;default:0;comment:排序标记"`
	Affix     bool   `gorm:"type:tinyint(1);not null;default:0;comment:固定地址栏"`
	Status    int32  `gorm:"type:tinyint(2);not null;default:1;comment:菜单状态 1-正常 2-停用"`
	CreatedAt int64  `gorm:"type:bigint(20);not null;default:0;comment:创建时间"`
	CreatedBy string `gorm:"type:varchar(128);not null;default:'';comment:创建人"`
	UpdatedAt int64  `gorm:"type:bigint(20);not null;default:0;comment:更新时间"`
	UpdatedBy string `gorm:"type:varchar(128);not null;default:'';comment:更新人"`
	IsDelete  bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否删除"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (SysMenu) TableComment() string {
	return "系统菜单表"
}

type SysApi struct {
	Id      int32  `gorm:"type:int Auto_Increment;primary_key;not null"`
	Title   string `gorm:"type:varchar(128);not null;default:'';comment:显示名称;index"`
	Path    string `gorm:"type:varchar(256);not null;default:'';comment:路由地址;index"`
	Action  string `gorm:"type:varchar(8);not null;default:'';comment:类型 GET POST PUT DELETE"`
	Desc    string `gorm:"type:varchar(256);not null;default:'';comment:描述"`
	NeedPer bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否需要权限"`
}

func (SysApi) TableName() string {
	return "sys_api"
}

func (SysApi) TableComment() string {
	return "系统菜单接口表"
}

type SysMenuApi struct {
	Id        int32 `gorm:"type:int Auto_Increment;primary_key;not null"`
	SysMenuId int32 `gorm:"type:int;not null;default:0;comment:菜单id"`
	SysApiId  int32 `gorm:"type:int;not null;default:0;comment:接口id"`
}

func (SysMenuApi) TableName() string {
	return "sys_menu_api"
}

func (SysMenuApi) TableComment() string {
	return "系统菜单接口关联表"
}

type SysCasbinRule struct {
	Id    int32  `gorm:"type:int Auto_Increment;primary_key;not null"`
	Ptype string `gorm:"type:varchar(100);not null;default:'';comment:策略类型"`
	V0    string `gorm:"type:varchar(100);not null;default:'';comment:角色"`
	V1    string `gorm:"type:varchar(100);not null;default:'';comment:路径"`
	V2    string `gorm:"type:varchar(100);not null;default:'';comment:请求方法"`
	V3    string `gorm:"type:varchar(100);not null;default:'';comment:保留字段"`
	V4    string `gorm:"type:varchar(100);not null;default:'';comment:保留字段"`
	V5    string `gorm:"type:varchar(100);not null;default:'';comment:保留字段"`
}

func (SysCasbinRule) TableName() string {
	return "sys_casbin_rule"
}

type SearchSysMenuFilter struct {
	MenuCode string
	MenuType int32
}

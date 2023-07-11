package model

import "github.com/spf13/cast"

type SysUser struct {
	UserId    int64  `gorm:"type:int(10) auto_increment;primary_key;用户id"`              // 用户ID
	Username  string `gorm:"type:varchar(64);NOT NULL;DEFAULT '';COMMENT:用户名"`         // 用户名
	Password  string `gorm:"type:varchar(255);NOT NULL;DEFAULT '';COMMENT:密码"`          // 密码
	NickName  string `gorm:"type:varchar(64);NOT NULL;DEFAULT '';COMMENT:昵称"`           // 昵称
	Avatar    string `gorm:"type:varchar(255);NOT NULL;DEFAULT '';COMMENT:头像"`          // 头像
	Phone     string `gorm:"type:varchar(11);NOT NULL;DEFAULT '';COMMENT:手机号"`         // 手机号
	Email     string `gorm:"type:varchar(128);NOT NULL;DEFAULT '';COMMENT:邮箱"`          // 邮箱
	Sex       string `gorm:"type:varchar(8);NOT NULL;DEFAULT '';COMMENT:性别"`            // 性别
	RoleId    int64  `gorm:"type:int(10);NOT NULL;DEFAULT 0;COMMENT:角色id"`              // 角色id
	DeptId    int64  `gorm:"type:int(10);NOT NULL;DEFAULT 0;COMMENT:部门id"`              // 部门id
	PostId    int64  `gorm:"type:int(10);NOT NULL;DEFAULT 0;COMMENT:岗位id"`              // 岗位id
	Remark    string `gorm:"type:varchar(255);NOT NULL;DEFAULT '';COMMENT:备注"`          // 备注
	Status    int8   `gorm:"type:tinyint(1);NOT NULL;DEFAULT 2;COMMENT:状态 1启用 2停用"` // 状态 1启用 2停用
	CreatedAt int64  `gorm:"type:bigint;NOT NULL;DEFAULT 0;COMMENT:创建时间"`             // 创建时间
	CreatedBy string `gorm:"type:varchar(128);NOT NULL;DEFAULT '';COMMENT:创建人"`        // 创建人
	UpdatedAt int64  `gorm:"type:bigint;NOT NULL;DEFAULT 0;COMMENT:更新时间"`             // 更新时间
	DeletedAt int64  `gorm:"type:bigint;NOT NULL;DEFAULT 0;COMMENT:删除时间"`             // 删除时间
}

func (SysUser) TableName() string {
	return "sys_user"
}

type UserMeta struct {
	UserId   int64
	UserName string
	NickName string
	RoleId   int64
}

func (u *UserMeta) Unique() string {
	return cast.ToString(u.UserId)
}

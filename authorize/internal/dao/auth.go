package dao

import (
	"github.com/team-dandelion/go-dandelion/application"
	"go-admin-example/authorize/internal/model"
)

type IAuth interface {
	GetUserInfoByUserName(userName string) (user model.SysUser, err error)
}

func NewAuth() IAuth {
	return &authDao{}
}

type authDao struct {
	application.DB
	application.Redis
}

func (d *authDao) GetUserInfoByUserName(userName string) (user model.SysUser, err error) {
	err = d.GetDB().Model(model.SysUser{}).Where("user_name = ?", userName).First(&user).Error
	return
}

package dao

import (
	"go-admin-example/authorize/internal/model"
	"gorm.io/gorm"
)

var Auth authDao

type authDao struct {
}

func (d *authDao) GetUserInfoByUserName(tx *gorm.DB, userName string) (user model.SysUser, err error) {
	err = tx.Model(model.SysUser{}).Where("username = ?", userName).First(&user).Error
	return
}

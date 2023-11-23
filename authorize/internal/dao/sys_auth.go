package dao

import (
	"github.com/team-dandelion/go-dandelion/application"
	"go-admin-example/authorize/internal/enum"
	"go-admin-example/authorize/internal/model"
)

type IAuth interface {
	GetUserInfoByUserName(userName string) (user model.SysUser, err error)
	GetUserMenus(filter model.UserMenusFilter) (menus []model.SysMenu, err error)
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

func (d *authDao) GetUserMenus(filter model.UserMenusFilter) (
	menus []model.SysMenu, err error) {
	if filter.IsSuper {
		err = d.GetDB().Model(model.SysMenu{}).
			Where("status = ?", enum.MenuStatusEnable).
			Where("is_delete = ?", false).
			Find(&menus).Error
		return
	}

	// 获取用户角色
	err = d.GetDB().Model(model.SysMenu{}).
		Joins("left join sys_role_menu on sys_menu.id = sys_role_menu.menu_id").
		Joins("left join sys_user_role on sys_role_menu.role_id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ?", filter.UserId).
		Where("sys_menu.status = ?", enum.MenuStatusEnable).
		Where("sys_menu.is_delete = ?", false).
		Find(&menus).Error
	return
}

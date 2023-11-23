package dao

import (
	"github.com/team-dandelion/go-dandelion/application"
	"go-admin-example/authorize/internal/model"
)

type ISysMenu interface {
	SearchSysMenu(model.SearchSysMenuFilter) ([]model.SysMenu, error)
}

func NewSysMenu() ISysMenu {
	return &sysMenuDao{}
}

type sysMenuDao struct {
	application.DB
	application.Redis
}

func (s sysMenuDao) SearchSysMenu(filter model.SearchSysMenuFilter) ([]model.SysMenu, error) {
	//TODO implement me
	panic("implement me")
}

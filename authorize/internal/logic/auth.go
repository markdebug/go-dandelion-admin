package logic

import (
	"github.com/gly-hub/dandelion-plugs/jwt"
	"github.com/team-dandelion/go-dandelion/logger"
	"go-admin-example/authorize/Tools"
	"go-admin-example/authorize/internal/dao"
	"go-admin-example/authorize/internal/enum"
	"go-admin-example/authorize/internal/model"
	authModel "go-admin-example/common/model/authorize"
	"gorm.io/gorm"
)

type IAuth interface {
	Login(params authModel.LoginParams) (string, error)
	Logout(opt model.CtxOption) error
	VerifyPermission(params authModel.VerifyPermissionParams) (meta model.UserMeta, err error)
	GetUserMenus(opt model.CtxOption) (menus []model.SysMenu, err error)
}

func NewAuth() IAuth {
	return &authLogic{
		AuthDao: dao.NewAuth(),
	}
}

type authLogic struct {
	AuthDao dao.IAuth
}

// Login 登录
func (l *authLogic) Login(params authModel.LoginParams) (string, error) {
	userInfo, uErr := l.AuthDao.GetUserInfoByUserName(params.UserName)
	if uErr != nil && uErr != gorm.ErrRecordNotFound {
		logger.Error(uErr)
		return "", enum.DataBaseError
	}

	// 判断信息
	params.Password = Tools.Md5V(params.Password)
	if uErr == gorm.ErrRecordNotFound || userInfo.Password != params.Password {
		return "", enum.UserNameOrPasswordError
	}

	// 生成token
	token, tErr := jwt.Token(&model.UserMeta{
		UserId:   userInfo.Id,
		UserName: userInfo.UserName,
		NickName: userInfo.Nickname,
		IsSuper:  userInfo.IsSuper,
	})

	if tErr != nil {
		logger.Error(tErr)
		return "", enum.UserNameOrPasswordError
	}

	return token, nil
}

// Logout 登出
func (l *authLogic) Logout(opt model.CtxOption) error {
	_ = jwt.Del(&model.UserMeta{
		UserId: opt.UserId,
	})
	return nil
}

// VerifyPermission 校验权限
func (l *authLogic) VerifyPermission(params authModel.VerifyPermissionParams) (
	meta model.UserMeta, err error) {
	err = jwt.Parse(params.Token, &meta)
	return
}

// GetUserMenus 获取用户菜单列表
func (l *authLogic) GetUserMenus(opt model.CtxOption) (menus []model.SysMenu, err error) {
	menus, err = l.AuthDao.GetUserMenus(model.UserMenusFilter{
		UserId:  opt.UserId,
		IsSuper: true,
	})
	return
}

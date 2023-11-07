package logic

import (
	"github.com/gly-hub/dandelion-plugs/captcha"
	"github.com/gly-hub/dandelion-plugs/jwt"
	"github.com/rs/xid"
	"github.com/team-dandelion/go-dandelion/logger"
	"go-admin-example/authorize/internal/dao"
	"go-admin-example/authorize/internal/enum"
	"go-admin-example/authorize/internal/model"
	"go-admin-example/common/model/authorize"
	"gorm.io/gorm"
)

type IAuth interface {
	Login(params authorize.LoginParams) (string, error)
	Logout(userId int64) error
	GenerateCaptcha() (content, id string, err error)
}

func NewAuth() IAuth {
	return &authLogic{
		AuthDao: dao.NewAuth(),
	}
}

type authLogic struct {
	AuthDao dao.IAuth
}

func (l *authLogic) Login(params authorize.LoginParams) (string, error) {
	userInfo, uErr := l.AuthDao.GetUserInfoByUserName(params.UserName)
	if uErr != nil && uErr != gorm.ErrRecordNotFound {
		logger.Error(uErr)
		return "", enum.DataBaseError
	}

	// 0判断信息
	if uErr == gorm.ErrRecordNotFound || userInfo.Password != params.Password {
		return "", enum.UserNameOrPasswordError
	}

	if !captcha.Verify(params.CaptchaId, params.CaptchaCode) {
		return "", enum.CaptchaError
	}

	// 生成token
	token, tErr := jwt.Token(&model.UserMeta{
		UserId:   userInfo.Id,
		UserName: userInfo.UserName,
		NickName: userInfo.Nickname,
	})

	if tErr != nil {
		logger.Error(tErr)
		return "", enum.UserNameOrPasswordError
	}

	return token, nil
}

func (l *authLogic) GenerateCaptcha() (content, id string, err error) {
	id = xid.New().String()
	img, _, cErr := captcha.Create(id)

	if cErr != nil {
		return "", id, cErr
	}
	content = img.Base64()
	return content, id, nil
}

func (l *authLogic) Logout(userId int64) error {
	_ = jwt.Del(&model.UserMeta{
		UserId: userId,
	})
	return nil
}

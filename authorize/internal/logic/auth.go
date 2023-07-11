package logic

import (
	"github.com/gly-hub/dandelion-plugs/captcha"
	"github.com/gly-hub/dandelion-plugs/jwt"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/rs/xid"
	"go-admin-example/authorize/internal/dao"
	"go-admin-example/authorize/internal/enum"
	"go-admin-example/authorize/internal/model"
	"go-admin-example/common/model/authorize"
	"gorm.io/gorm"
)

var Auth authLogic

type authLogic struct {
	application.DB
	application.Redis
}

func (l *authLogic) Login(params authorize.LoginParams) (string, error) {
	userInfo, uErr := dao.Auth.GetUserInfoByUserName(l.GetDB(), params.UserName)
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

	// 生成token TODO
	token, tErr := jwt.Token(&model.UserMeta{
		UserId:   userInfo.UserId,
		UserName: userInfo.Username,
		NickName: userInfo.NickName,
		RoleId:   userInfo.RoleId,
	})

	if tErr != nil {
		logger.Error(tErr)
		return "", enum.UserNameOrPasswordError
	}

	return token, nil
}

func (l *authLogic) GenerateCaptcha() (content, id string, err error) {
	id = xid.New().String()
	img, cErr := captcha.Create(id)
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

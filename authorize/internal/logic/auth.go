package logic

import (
	"fmt"
	"github.com/gly-hub/dandelion-plugs/captcha"
	"github.com/rs/xid"
	"go-admin-example/common/model/authorize"
)

var Auth authLogic

type authLogic struct {
}

func (l *authLogic) Login(params authorize.LoginParams) error {
	ok := captcha.Verify(params.CaptchaId, params.CaptchaCode)
	fmt.Println(ok)
	return nil
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

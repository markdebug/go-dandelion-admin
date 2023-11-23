package logic

import (
	"github.com/gly-hub/dandelion-plugs/captcha"
	"github.com/gly-hub/dandelion-plugs/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/config"
	"go-admin-example/authorize/internal/dao"
	"go-admin-example/authorize/internal/enum"
	"go-admin-example/authorize/internal/model"
	"go-admin-example/common/model/authorize"
	"gorm.io/gorm"
	"testing"
)

func NewMockAuth() dao.IAuth {
	return &authMockDao{}
}

type authMockDao struct {
}

func (d *authMockDao) GetUserInfoByUserName(userName string) (user model.SysUser, err error) {
	if userName == "" {
		err = gorm.ErrRecordNotFound
		return
	}

	user = model.SysUser{
		Id:       100,
		UserName: "1111",
		Password: "123456",
	}

	return
}

func InitApplication() {
	config.InitConfig("local")
	application.Init()
}

func Test_authLogic_Login(t *testing.T) {
	InitApplication()
	_ = application.Plugs(captcha.Plug(), jwt.Plug())

	logicAuth := authLogic{AuthDao: NewMockAuth()}
	params := authorize.LoginParams{
		UserName: "11111",
	}
	token, err := logicAuth.Login(params.UserName, params.Password)
	assert.Equal(t, err, enum.UserNameOrPasswordError)
	assert.Empty(t, token)

	params.UserName = "1111"
	params.Password = "123"
	token, err = logicAuth.Login(params.UserName, params.Password)
	assert.Equal(t, err, enum.UserNameOrPasswordError)
	assert.Empty(t, token)
}

func Test_authLogic_Logout(t *testing.T) {
	InitApplication()
	_ = application.Plugs(jwt.Plug())
	logicAuth := NewAuth()
	err := logicAuth.Logout(100)
	assert.Nil(t, err)
}

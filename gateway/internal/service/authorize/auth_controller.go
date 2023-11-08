package authorize

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/server/http"
	authModel "go-admin-example/common/model/authorize"
	rpcService "go-admin-example/common/service"
)

type AuthController struct {
	http.HttpController
}

// Login
// @Summary 用户登录
// @Description 用户登录api接口
// @Tags 基础模块
// @Param deptName body auth.LoginParams true "登录参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/login [post]
func (a *AuthController) Login(c *routing.Context) error {
	return application.SRpcCall(c, rpcService.AuthorizeService, rpcService.AuthorizeFuncLogin, new(authModel.LoginParams), new(authModel.LoginResp))
}

// Captcha
// @Summary 验证码获取
// @Description 验证码获取api接口
// @Tags 基础模块
// @Param deptName body authModel.CaptchaParams true "参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/captcha [get]
func (a *AuthController) Captcha(c *routing.Context) error {
	var (
		req  = new(authModel.CaptchaParams)
		resp = new(authModel.CaptchaResp)
	)
	err := application.RpcCall(c, rpcService.AuthorizeService, rpcService.AuthorizeFuncGenerateCaptcha, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}

// Logout
// @Summary 注销登录
// @Description 注销登录api接口
// @Tags 基础模块
// @Param deptName body authModel.LoginParams true "参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/logout [delete]
func (a *AuthController) Logout(c *routing.Context) error {
	var (
		req  = new(authModel.LoginParams)
		resp = new(authModel.LoginResp)
	)
	err := application.RpcCall(c, rpcService.AuthorizeService, rpcService.AuthorizeFuncGenerateCaptcha, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}

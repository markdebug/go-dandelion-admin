package service

import (
	"context"
	errorx "github.com/gly-hub/go-dandelion/error-support"
	"go-admin-example/authorize/internal/logic"
	"go-admin-example/common/model/authorize"
)

func (s *RpcApi) Login(ctx context.Context, req authorize.LoginParams, resp *authorize.LoginResp) error {
	err := logic.Auth.Login(req)
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	return nil
}

// GenerateCaptcha 生成验证码
func (s *RpcApi) GenerateCaptcha(ctx context.Context, req authorize.CaptchaParams, resp *authorize.CaptchaResp) error {
	content, id, err := logic.Auth.GenerateCaptcha()
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	resp.Id = id
	resp.Img = content
	return nil
}

package service

import (
	"context"
	errorx "github.com/team-dandelion/go-dandelion/error-support"
	"go-admin-example/common/model/authorize"
)

func (s *RpcApi) Login(ctx context.Context, req authorize.LoginParams, resp *authorize.LoginResp) error {
	_, err := s.AuthLogic.Login(req)
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	return nil
}

func (s *RpcApi) Logout(ctx context.Context, req authorize.LoginParams, resp *authorize.LoginResp) error {
	err := s.AuthLogic.Logout(s.UserId(ctx))
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	return nil
}

// GenerateCaptcha 生成验证码
func (s *RpcApi) GenerateCaptcha(ctx context.Context, req authorize.CaptchaParams, resp *authorize.CaptchaResp) error {
	content, id, err := s.AuthLogic.GenerateCaptcha()
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	resp.Id = id
	resp.Img = content
	return nil
}

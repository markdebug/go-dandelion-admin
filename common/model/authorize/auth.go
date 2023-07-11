package authorize

import "go-admin-example/common/model/lib"

type (
	CaptchaParams struct {
	}

	CaptchaResp struct {
		lib.Response
		Img string `json:"img"`
		Id  string `json:"id"`
	}
)

type (
	LoginParams struct {
		UserName    string `json:"user_name"`    // 用户名
		Password    string `json:"password"`     // 登录密码
		CaptchaCode string `json:"captcha_code"` // 验证码
		CaptchaId   string `json:"captcha_id"`   // 验证码id
	}

	LoginResp struct {
		Token string `json:"token"`
		lib.Response
	}
)

type (
	LogoutParams struct {
	}

	LogoutResp struct {
		lib.Response
	}
)

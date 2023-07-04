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
		CaptchaCode string `json:"captcha_code"`
		CaptchaId   string `json:"captcha_id"`
	}

	LoginResp struct {
		lib.Response
	}
)

package boot

import (
	"github.com/gly-hub/dandelion-plugs/captcha"
	"github.com/gly-hub/go-dandelion/application"
)

func Init() {
	// 将需要初始化的方法在该处注册
	_ = application.Plugs(captcha.Plug())
}

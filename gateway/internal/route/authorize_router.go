package route

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"go-admin-example/gateway/internal/service/authorize"
)

func initAuthRoute(baseRouter *routing.RouteGroup) {
	authHandler := authorize.AuthController{}
	// 登录登出
	baseRouter.Post("/login", authHandler.Login)
	baseRouter.Get("/captcha", authHandler.Captcha)
}

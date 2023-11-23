package route

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/team-dandelion/analysis-plug"
	"go-admin-example/gateway/internal/middleware"
	"go-admin-example/gateway/internal/service/authorize"
)

func initAuthRoute(baseRouter *routing.RouteGroup) {
	authHandler := authorize.AuthController{}
	// 登录登出
	baseRouter.Use(analysis.HttpPrometheus())
	baseRouter.Use(middleware.PermissionMiddleware())
	baseRouter.Post("/login", authHandler.Login)
	baseRouter.Delete("/logout", authHandler.Logout)
	baseRouter.Get("/user_menu", authHandler.UserMenus)
}

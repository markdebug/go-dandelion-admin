package middleware

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/server/http"
	authModel "go-admin-example/common/model/authorize"
	rpcService "go-admin-example/common/service"
)

func PermissionMiddleware() routing.Handler {
	return func(c *routing.Context) error {
		// 判断是否忽略权限校验 TODO
		if IgnoreUrl(string(c.Path())) {
			return c.Next()
		}

		var resp = new(authModel.VerifyPermissionResp)
		if err := application.RpcCall(c, rpcService.AuthorizeService,
			rpcService.AuthorizeFuncVerifyPermission,
			authModel.VerifyPermissionParams{
				Token:  c.Header.Value("token"),
				Path:   string(c.Path()),
				Method: string(c.Method()),
			}, resp); err != nil {
			return (&http.HttpController{}).Fail(c, err)
		}
		c.Header.SetInt32("UserId", resp.UserId)
		c.Header.Set("UserName", resp.UserName)
		c.Header.SetBool("IsSuper", resp.IsSuper)
		c.Header.Set("Nickname", resp.NickName)

		return c.Next()
	}
}

func IgnoreUrl(path string) bool {
	urls := []string{
		"/api/login",
	}

	for _, url := range urls {
		if url == path {
			return true
		}
	}
	return false
}

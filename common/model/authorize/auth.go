package authorize

import "go-admin-example/common/model/lib"

type (
	LoginParams struct {
		UserName string `json:"username"` // 用户名
		Password string `json:"password"` // 登录密码
	}

	LoginResp struct {
		Token    string   `json:"token"`
		UserInfo UserInfo `json:"userInfo"`
		lib.Response
	}

	UserInfo struct {
		UserId    string   `json:"userId"`
		UserName  string   `json:"userName"`
		Role      []string `json:"role"`
		Dashboard string   `json:"dashboard"`
	}
)

type (
	LogoutParams struct {
	}

	LogoutResp struct {
		lib.Response
	}
)

type (
	VerifyPermissionParams struct {
		Token  string
		Path   string
		Method string
	}

	VerifyPermissionResp struct {
		lib.Response
		UserName string
		UserId   int32
		NickName string
		IsSuper  bool
	}
)

type (
	GetUserMenusParams struct {
	}

	GetUserMenusResp struct {
		lib.Response
		DashboardGrid []string   `json:"dashboardGrid"`
		Menu          []UserMenu `json:"menu"`
		Permissions   []string   `json:"permissions"`
	}

	UserMenu struct {
		Name      string       `json:"name"`
		Path      string       `json:"path"`
		Sort      int32        `json:"sort"`
		Meta      UserMenuMeta `json:"meta"`
		Component string       `json:"component"`
		Children  []UserMenu   `json:"children"`
	}

	UserMenuMeta struct {
		Icon  string `json:"icon"`
		Title string `json:"title"`
		Type  string `json:"type"`
		Tag   string `json:"tag"`
		Affix bool   `json:"affix"`
	}
)

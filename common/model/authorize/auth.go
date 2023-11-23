package authorize

import "go-admin-example/common/model/lib"

type (
	LoginParams struct {
		UserName string `json:"user_name"` // 用户名
		Password string `json:"password"`  // 登录密码
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
		Menus []UserMenu `json:"menus"`
	}

	UserMenu struct {
		Id        int32      `json:"id"`
		MenuType  int32      `json:"menu_type"`
		MenuCode  string     `json:"menu_code"`
		Title     string     `json:"title"`
		Icon      string     `json:"icon"`
		Component string     `json:"component"`
		Path      string     `json:"path"`
		Sort      int32      `json:"sort"`
		Children  []UserMenu `json:"children"`
	}
)

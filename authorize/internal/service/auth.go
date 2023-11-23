package service

import (
	"context"
	errorx "github.com/team-dandelion/go-dandelion/error-support"
	"go-admin-example/authorize/Tools/customesort"
	"go-admin-example/authorize/internal/model"
	"go-admin-example/common/model/authorize"
	"sort"
)

func (s *RpcApi) Login(ctx context.Context, req authorize.LoginParams, resp *authorize.LoginResp) (err error) {
	resp.Token, err = s.AuthLogic.Login(req)
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	return nil
}

func (s *RpcApi) Logout(ctx context.Context, req authorize.LogoutParams, resp *authorize.LogoutResp) error {
	err := s.AuthLogic.Logout(s.CtxOption(ctx))
	if err != nil {
		errorx.Format(err, resp)
		return nil
	}
	return nil
}

func (s *RpcApi) VerifyPermission(ctx context.Context, req authorize.VerifyPermissionParams, resp *authorize.VerifyPermissionResp) error {
	meta, mErr := s.AuthLogic.VerifyPermission(req)
	if mErr != nil {
		errorx.Format(mErr, resp)
		return nil
	}

	resp.UserId = meta.UserId
	resp.UserName = meta.UserName
	resp.NickName = meta.NickName
	resp.IsSuper = meta.IsSuper

	return nil
}

func (s *RpcApi) GetUserMenus(ctx context.Context, req authorize.GetUserMenusParams, resp *authorize.GetUserMenusResp) error {
	menus, mErr := s.AuthLogic.GetUserMenus(s.CtxOption(ctx))
	if mErr != nil {
		errorx.Format(mErr, resp)
		return nil
	}

	// 构建结果
	resp.Menus = s.buildMenuTree(0, menus)
	return nil
}

func (s *RpcApi) buildMenuTree(parentId int32, menus []model.SysMenu) (tree []authorize.UserMenu) {
	tree = make([]authorize.UserMenu, 0)
	for _, m := range menus {
		if m.ParentId == parentId {
			tree = append(tree, authorize.UserMenu{
				Id:        m.Id,
				MenuType:  m.MenuType,
				MenuCode:  m.Code,
				Title:     m.Title,
				Icon:      m.Icon,
				Component: m.Component,
				Path:      m.Path,
				Sort:      m.Sort,
				Children:  s.buildMenuTree(m.Id, menus),
			})
		}
	}
	// 排序
	sort.Sort(customesort.UserMenuSlice(tree))
	return
}

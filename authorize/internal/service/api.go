package service

import (
	"context"
	"github.com/team-dandelion/go-dandelion/server/rpcx"
	"go-admin-example/authorize/internal/logic"
	"go-admin-example/authorize/internal/model"
)

type RpcApi struct {
	AuthLogic logic.IAuth
}

func NewRpcApi() *RpcApi {
	return &RpcApi{
		AuthLogic: logic.NewAuth(),
	}
}

func (s *RpcApi) CtxOption(ctx context.Context) model.CtxOption {
	return model.CtxOption{
		UserName: rpcx.Header().Value(ctx, "UserName"),
		UserId:   rpcx.Header().Int32Default(ctx, "UserId", 0),
		IsSuper:  rpcx.Header().BoolDefault(ctx, "IsSuper", false),
	}
}

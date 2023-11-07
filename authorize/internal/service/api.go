package service

import (
	"context"
	"github.com/team-dandelion/go-dandelion/server/rpcx"
	"go-admin-example/authorize/internal/logic"
)

type RpcApi struct {
	AuthLogic logic.IAuth
}

func NewRpcApi() *RpcApi {
	return &RpcApi{
		AuthLogic: logic.NewAuth(),
	}
}

func (s *RpcApi) UserId(ctx context.Context) int64 {
	return rpcx.Header().Int64Default(ctx, "user_id", 0)
}

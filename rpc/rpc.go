package rpc

import (
	"context"
	lkCommon "github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/protos"
	"github.com/topfreegames/pitaya"
)

// SendRpc sends rpc
func SendRpc(ctx context.Context, rpcMsg *protos.RpcMsg) (*protos.RpcRes, *protos.LResult) {
	ret := &protos.RpcRes{}
	err := pitaya.RPCTo(ctx, rpcMsg.ServerId, rpcMsg.Route, ret, rpcMsg)
	if err != nil {
		return nil, lkCommon.OfFailData(lkCommon.RetSystemFail, lkCommon.CodeRpcFail, "RPC-000", err.Error())
	}
	return ret, lkCommon.Ok()
}

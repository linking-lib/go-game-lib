package common

import (
	"encoding/json"
	"github.com/linking-lib/go-game-lib/protos"
)

func NewRpcMsg(route string, Msg string, token string) *protos.RpcMsg {
	return &protos.RpcMsg{Route: route, ServerId: "", Msg: Msg, Token: token}
}

func OkRpc() *protos.RpcRes {
	return &protos.RpcRes{Code: RetOk, ErrCode: CodeOk, Msg: MsgOk}
}

func OkRpcData(data interface{}) *protos.RpcRes {
	result := OkRpc()
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

func FailRpc(code string, errCode string, msg string) *protos.RpcRes {
	return &protos.RpcRes{Code: code, ErrCode: errCode, Msg: msg}
}

func FailRpcData(code string, errCode string, msg string, data interface{}) *protos.RpcRes {
	result := FailRpc(code, errCode, msg)
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

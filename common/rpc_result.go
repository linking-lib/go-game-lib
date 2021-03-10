package common

import (
	"encoding/json"
	"github.com/linking-lib/go-game-lib/protos"
)

type RpcMsg struct {
	Route    string `json:"route,omitempty"`
	ServerId string `json:"server_id,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Token    string `json:"token,omitempty"`
}

type RpcResult struct {
	Code    string      `json:"code,omitempty"`
	ErrCode string      `json:"err_code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

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

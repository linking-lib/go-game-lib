package common

import (
	"encoding/json"
	"github.com/linking-lib/go-game-lib/protos"
)

type LResult struct {
	Api     string      `json:"api,omitempty"`
	Code    int32       `json:"code,omitempty"`
	ErrCode string      `json:"err_code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func TestFail(lResult *protos.LResult) bool {
	return lResult.Code != RetOk
}

func Ok() *protos.LResult {
	return &protos.LResult{Code: RetOk, ErrCode: CodeOk, Msg: MsgOk}
}

func OfOk(api string) *protos.LResult {
	return &protos.LResult{Api: api, Code: RetOk, ErrCode: CodeOk, Msg: MsgOk}
}

func OfOkData(api string, data interface{}) *protos.LResult {
	result := OfOk(api)
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

func OfFail(api string, code int32, errCode string, msg string) *protos.LResult {
	return &protos.LResult{Api: api, Code: code, ErrCode: errCode, Msg: msg}
}

func OfFailData(api string, code int32, errCode string, msg string, data interface{}) *protos.LResult {
	result := OfFail(api, code, errCode, msg)
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

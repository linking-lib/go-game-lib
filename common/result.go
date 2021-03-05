package common

import (
	"encoding/json"
	"github.com/linking-lib/go-game-lib/protos"
)

type LResult struct {
	Api     string      `json:"api,omitempty"`
	Code    int32       `json:"code"`
	ErrCode string      `json:"err_code"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func TestFail(lResult *protos.LResult) bool {
	return lResult.Code != RetOk
}

func Ok() *protos.LResult {
	return &protos.LResult{Code: RetOk, ErrCode: CodeOk, Msg: MsgOk}
}

func OfOkData(data interface{}) *protos.LResult {
	result := Ok()
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

func OfFail(code int32, errCode string, msg string) *protos.LResult {
	return &protos.LResult{Code: code, ErrCode: errCode, Msg: msg}
}

func OfFailData(code int32, errCode string, msg string, data interface{}) *protos.LResult {
	result := OfFail(code, errCode, msg)
	if data != nil {
		jsonB, _ := json.Marshal(data)
		result.Data = string(jsonB)
	}
	return result
}

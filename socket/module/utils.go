package module

import (
	"context"
	"github.com/alecthomas/log4go"
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/lkerrors"
	"github.com/linking-lib/go-game-lib/protos"
	manager "github.com/linking-lib/go-game-lib/utils"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/util"
	"reflect"
)

type ApiProcessMode string

const (
	ApiModeMain   ApiProcessMode = "Main"
	ApiModeThread ApiProcessMode = "Thread"
	ApiModeNone   ApiProcessMode = "None"
)

type HandlerMsg struct {
	Uid     string
	ApiType ApiProcessMode
	Api     string
	Msg     interface{}
}

func DoHandleMsg(msg HandlerMsg) (*protos.LResult, error) {
	handler := handlers[msg.Api]
	args := []reflect.Value{handler.Receiver, reflect.ValueOf(msg)}
	ret, err := util.Pcall(handler.Method, args)
	if err != nil {
		_ = log4go.Error("DoHandleMsg=============" + err.Error())
		return common.OfFail(msg.Api, lkerrors.RetInternalCode, lkerrors.ErrInternalCode, err.Error()), err
	}
	if ret != nil {
		return ret.(*protos.LResult), nil
	}
	return common.Ok(), nil
}

// SendRPC sends rpc
func SendRPC(ctx context.Context, serverId string, route string, request *protos.LRequest) {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	ret := &protos.LResult{}
	err := pitaya.RPCTo(ctx, serverId, route, ret, request)
	if err != nil {
		logger.Errorf("Failed to execute RPCTo %s - %s", route, err.Error())
	}
}

/**
加入组
*/
func JoinGroup(ctx context.Context, group string, uid string) bool {
	logger := manager.GetLog(ctx)
	// 1、判断用户是否已经在组中，如果不存在再加入
	flag, err := pitaya.GroupContainsMember(ctx, group, uid)
	if err != nil {
		logger.Error("Failed to contains group member: " + err.Error())
		return false
	}
	if !flag {
		err = pitaya.GroupAddMember(ctx, group, uid)
		if err != nil {
			logger.Error("Failed to join group: " + err.Error())
			return false
		}
	}
	return true
}

/**
离开组
*/
func LeaveGroup(ctx context.Context, group string, uid string) bool {
	logger := manager.GetLog(ctx)
	// 1、用户从组中移除
	err := pitaya.GroupRemoveMember(ctx, group, uid)
	if err != nil {
		logger.Error("Failed to leave group member: " + err.Error())
		return false
	}
	return true
}

/**
清空组
*/
func ClearGroup(ctx context.Context, group string) bool {
	logger := manager.GetLog(ctx)
	// 1、清空组中成员
	err := pitaya.GroupRemoveAll(ctx, group)
	if err != nil {
		logger.Error("Failed to clear group: " + err.Error())
		return false
	}
	return true
}

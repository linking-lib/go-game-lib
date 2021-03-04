package module

import (
	"context"
	"github.com/linking-lib/go-game-lib/socket/session"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/groups"
)

type SelfComponent interface {
	component.Component
}

// Base implements a default module for Component.
type SelfBase struct {
	component.Base
}

type GroupMsg struct {
	Group string
	Uid   string
	Msg   string
	Param map[string]interface{}
}

func (g *GroupMsg) GroupParam(key string, value interface{}) {
	if g.Param == nil {
		g.Param = make(map[string]interface{})
	}
	g.Param[key] = value
}

func (b *SelfBase) InitGroup(conf *config.Config) {
	var gsi groups.GroupService
	var err error
	if conf != nil {
		gsi, err = groups.NewEtcdGroupService(conf, nil)
	} else {
		gsi = groups.NewMemoryGroupService(conf)
	}
	if err != nil {
		panic(err)
	}
	// 初始组
	pitaya.InitGroups(gsi)
}

func (b *SelfBase) CreateGroup(group string) {
	// 创建组
	_ = pitaya.GroupCreate(context.Background(), group)
}

/**
绑定用户
*/
func (b *SelfBase) BindUser(ctx context.Context, uid string) error {
	// 1、从ctx中获得session
	s := session.GetSession(ctx)
	// 2、绑定session用户编号
	err := s.Bind(ctx, uid)
	if err != nil {
		return pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}
	return nil
}

// SetSessionData sets the session data
func (b *SelfBase) SetSessionOneData(ctx context.Context, key string, value interface{}) bool {
	data := make(map[string]interface{})
	data[key] = value
	return b.SetSessionData(ctx, data)
}

// SetSessionData sets the session data
func (b *SelfBase) SetSessionData(ctx context.Context, data map[string]interface{}) bool {
	logger := pitaya.GetDefaultLoggerFromCtx(ctx)
	s := pitaya.GetSessionFromCtx(ctx)
	// 1、取出老数据
	curData := s.GetData()
	// 2、将新值赋予老数据
	for k, v := range data {
		curData[k] = v
	}
	err := s.SetData(curData)
	if err != nil {
		logger.Error("Failed to set session data")
		logger.Error(err)
		return false
	}
	err = s.PushToFront(ctx)
	if err != nil {
		logger.Error("Failed to PushToFront session data")
		logger.Error(err)
		return false
	}
	return true
}

// GetSessionData gets the session data
func (b *SelfBase) GetSessionData(ctx context.Context, key string) interface{} {
	s := pitaya.GetSessionFromCtx(ctx)
	return s.Get(key)
}

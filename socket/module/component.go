package module

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/config"
	"github.com/topfreegames/pitaya/groups"
)

type SelfComponent interface {
	component.Component
	Group() string
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

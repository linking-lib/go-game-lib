package session

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/session"
)

// SessionData struct
type LkSessionData struct {
	Data map[string]interface{}
}

type LkSession struct {
}

// GetSession gets the session data
func GetSession(ctx context.Context) *session.Session {
	return pitaya.GetSessionFromCtx(ctx)
}

// GetData gets the session data
func (s *LkSession) GetData(ctx context.Context, key string) interface{} {
	lkSession := GetSession(ctx)
	return lkSession.Get(key)
}

// GetAllData gets the session data
func (s *LkSession) GetAllData(ctx context.Context) *LkSessionData {
	lkSession := GetSession(ctx)
	res := &LkSessionData{
		Data: lkSession.GetData(),
	}
	return res
}

// SetData sets the session data
func (s *LkSession) SetData(ctx context.Context, key string, value interface{}) (bool, error) {
	lkSession := GetSession(ctx)
	curData := lkSession.GetData()
	curData[key] = value
	err := lkSession.SetData(curData)
	if err != nil {
		return false, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	err = lkSession.PushToFront(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetAnyData sets the session data
func (s *LkSession) SetAnyData(ctx context.Context, data map[string]interface{}) (bool, error) {
	lkSession := GetSession(ctx)
	// 1、取出老数据
	curData := lkSession.GetData()
	// 2、将新值赋予老数据
	for k, v := range data {
		curData[k] = v
	}
	err := lkSession.SetData(curData)
	if err != nil {
		return false, pitaya.Error(err, "CN-000", map[string]string{"failed": "set any data"})
	}
	err = lkSession.PushToFront(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetAllData sets the session data
func (s *LkSession) SetAllData(ctx context.Context, data *LkSessionData) (bool, error) {
	lkSession := GetSession(ctx)
	err := lkSession.SetData(data.Data)
	if err != nil {
		return false, pitaya.Error(err, "CN-000", map[string]string{"failed": "set all data"})
	}
	err = lkSession.PushToFront(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NotifySessionData sets the session data
func (s *LkSession) NotifySessionData(ctx context.Context, data *LkSessionData) {
	lkSession := GetSession(ctx)
	err := lkSession.SetData(data.Data)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
	err = lkSession.PushToFront(ctx)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
}

/**
绑定用户
*/
func (s *LkSession) BindUser(ctx context.Context, uid string) error {
	// 1、从ctx中获得session
	lkSession := GetSession(ctx)
	// 2、绑定session用户编号
	err := lkSession.Bind(ctx, uid)
	if err != nil {
		return pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}
	return nil
}

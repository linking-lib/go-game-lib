package session

import (
	"context"
	"fmt"
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/protos"
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

// GetSessionData gets the session data
func (c *LkSession) GetSessionData(ctx context.Context) (*LkSessionData, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	res := &LkSessionData{
		Data: s.GetData(),
	}
	return res, nil
}

// SetSessionData sets the session data
func (c *LkSession) SetSessionData(ctx context.Context, data *LkSessionData) (*protos.LResult, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		return nil, pitaya.Error(err, "CN-000", map[string]string{"failed": "set data"})
	}
	err = s.PushToFront(ctx)
	if err != nil {
		return nil, err
	}
	return common.Ok(), nil
}

// NotifySessionData sets the session data
func (c *LkSession) NotifySessionData(ctx context.Context, data *LkSessionData) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.SetData(data.Data)
	if err != nil {
		fmt.Println("got error on notify", err)
	}
}

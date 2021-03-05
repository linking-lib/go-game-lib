package socket

import (
	"context"
	"fmt"
	"github.com/linking-lib/go-game-lib/linking"
	"github.com/linking-lib/go-game-lib/socket/module"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/route"
	"net/http"
	"strings"
)

func InitComponent(local module.SelfComponent, remote module.SelfComponent) {
	pitaya.Register(local,
		component.WithName(local.Group()),
		component.WithNameFunc(strings.ToLower),
	)
	if remote != nil {
		pitaya.RegisterRemote(remote,
			component.WithName(remote.Group()),
			component.WithNameFunc(strings.ToLower),
		)
	}
}

func ConfigureFrontend(port int, httpPort int, dictionary ...string) {
	num := len(linking.GetRoutes())
	for i := 0; i < num; i++ {
		AddRoute(linking.GetRoutes()[i])
	}
	var dict = make(map[string]uint16, 0)
	num = len(dictionary)
	if num > 0 {
		for i := 0; i < num; i++ {
			dict[dictionary[i]] = uint16(i)
		}
	}
	err := pitaya.SetDictionary(dict)

	if err != nil {
		fmt.Printf("error setting route dictionary %s\n", err.Error())
		panic(err)
	}
	StartListen(port, httpPort)
}

func StartListen(port int, httpPort int) {
	wsPort := fmt.Sprintf(":%d", port)
	tcp := acceptor.NewWSAcceptor(wsPort)
	pitaya.AddAcceptor(tcp)
	wsHttpPort := fmt.Sprintf(":%d", httpPort)
	go http.ListenAndServe(wsHttpPort, nil)
}

/**
添加服务路由
*/
func AddRoute(serverType string) {
	err := pitaya.AddRoute(serverType, func(
		ctx context.Context,
		route *route.Route,
		payload []byte,
		servers map[string]*cluster.Server,
	) (*cluster.Server, error) {
		for k := range servers {
			return servers[k], nil
		}
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
}

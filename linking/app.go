package linking

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"github.com/linking-lib/go-game-lib/common"
	"github.com/linking-lib/go-game-lib/infrastructure/rocketmq"
	"github.com/linking-lib/go-game-lib/socket/module"
	"github.com/linking-lib/go-game-lib/utils/strs"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
)

// App is the base linking struct
type App struct {
	isFrontend  bool
	frontend    string
	config      *viper.Viper
	configured  bool
	debug       bool
	routes      []string
	clientRoute string
}

var (
	app = &App{
		frontend:    "connector",
		clientRoute: "onMessage",
		debug:       false,
		configured:  false,
	}
)

// Configure configures the linking
func Configure(
	isFrontend bool,
	frontend string,
	clientRoute string,
	routes []string,
	config *viper.Viper,
	debug bool,
) {
	if app.configured {
		logger.Log.Warn("lk socket configured twice!")
	}
	app.isFrontend = isFrontend
	app.frontend = frontend
	app.clientRoute = clientRoute
	app.routes = routes
	app.config = config
	app.configured = true
	app.debug = debug
}

func IsFrontend() bool {
	return app.isFrontend
}

func GetFrontend() string {
	return app.frontend
}

func GetClientRoute() string {
	return app.clientRoute
}

func GetRoutes() []string {
	return app.routes
}

func IsDebug() bool {
	return app.debug
}

func HandleMsg(msg *module.HandlerMsg) {
	_, err := module.DoHandleMsg(*msg)
	if err != nil {
		b, _ := json.Marshal(msg)
		_ = log4go.Error("HandleMsg error=========msg:"+string(b), err)
	}
}

func SendUserMsg(uid string, v interface{}) error {
	return SendUsersMsg([]string{uid}, v)
}

func SendUsersMsg(uids []string, v interface{}) error {
	errUids, err := pitaya.SendPushToUsers(app.clientRoute, v, uids, app.frontend)
	if err != nil {
		logger.Log.Errorf("SendUserMsg error, UID=%d, Error=%s", errUids, err.Error())
		return err
	}
	return nil
}

func SendRocketMqMsg(name string, data interface{}, tag string, prop map[string]string) error {
	body := common.ConvertJson(data)
	id, _ := common.Snow.GetSnowflakeId()
	key := strs.Int64ToStr(id)
	_, err := rocketmq.PublishMsg(name, rocketmq.MqMsg{MessageBody: body, MessageKey: key, MessageTag: tag, Properties: prop})
	if err != nil {
		propStr := ""
		if prop != nil {
			propStr = common.ConvertJson(prop)
		}
		logger.Log.Errorf("SendRocketMqMsg error, name=%s, tag=%s, body=%s, prop=%s, Error=%s", name, tag, body, propStr, err.Error())
	}
	return err
}

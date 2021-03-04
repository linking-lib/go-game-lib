package common

import (
	"github.com/linking-lib/go-game-lib/infrastructure/redis"
	"github.com/linking-lib/go-game-lib/utils/strs"
)

type Runtimes struct {
}

func (*Runtimes) GetProtocolType() string {
	var dataType = redis.RGet("default", "runtime:protocol_type")
	if strs.IsEmpty(dataType) {
		return ProtocolJson.String()
	} else {
		return dataType
	}
}

func (*Runtimes) SetProtocolType(protocolType string) {
	redis.RSet("default", "runtime:protocol_type", protocolType)
}

var SelfRuntime = &Runtimes{}

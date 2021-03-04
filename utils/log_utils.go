package manager

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
)

func GetLog(ctx context.Context) logger.Logger {
	return pitaya.GetDefaultLoggerFromCtx(ctx)
}

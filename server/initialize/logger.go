package initialize

import (
	"go.uber.org/zap"
)

func Logger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

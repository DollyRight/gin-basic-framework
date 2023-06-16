package initialize

import (
	"gin-basic-framework/global"
	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	mode := global.GLOBAL_VP.GetString("setting.mode")
	var logger *zap.Logger
	var err error
	if mode == "develop" {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	} else if mode == "production" {
		logger, err = zap.NewProduction()
		if err != nil {
			// 处理错误
			panic(err)
		}
	}

	defer logger.Sync()
	return logger
}

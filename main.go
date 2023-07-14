package main

import (
	"gin-basic-framework/core"
	"gin-basic-framework/global"
	"gin-basic-framework/initialize"
)

func main() {
	global.GLOBAL_VP = initialize.Viper()
	global.GLOBAL_LOGGER = initialize.Zap()
	global.GLOBAL_DB = initialize.Gorm()
	global.GLOBAL_NSQ_CONFIG = initialize.NsqConfig()
	global.GLOBAL_NSQ_PRODUCER = initialize.Producer()
	global.GLOBAL_NSQ_CONSUMER_POOL = initialize.ConsumerPool()
	if global.GLOBAL_DB != nil {
		initialize.RegisterTables()
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}
	core.StartServer()
}

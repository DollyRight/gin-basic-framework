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
	if global.GLOBAL_DB != nil {
		initialize.RegisterTables()
		db, _ := global.GLOBAL_DB.DB()
		defer db.Close()
	}
	core.StartServer()
}

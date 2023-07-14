package initialize

import (
	"gin-basic-framework/global"
	"gin-basic-framework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dsn := global.GLOBAL_CONFIG.DB.DsnProvider()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic("error in mysql connection")
	}
	return db
}

func RegisterTables() {
	db := global.GLOBAL_DB
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		global.GLOBAL_LOGGER.Fatal("register table failed")
	}
	global.GLOBAL_LOGGER.Info("register table suceess")
}

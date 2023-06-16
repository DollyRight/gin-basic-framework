package config

import (
	"fmt"
)

type DB struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

func (db *DB) DsnProvider() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
	)
	return dsn
}

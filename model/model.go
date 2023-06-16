package model

import (
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Model struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

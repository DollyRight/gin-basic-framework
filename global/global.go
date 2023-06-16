package global

import (
	"gin-basic-framework/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GLOBAL_VP     *viper.Viper
	GLOBAL_CONFIG config.Config
	GLOBAL_LOGGER *zap.Logger
	GLOBAL_DB     *gorm.DB
	BlackCache    local_cache.Cache
	lock          sync.RWMutex
)

package global

import (
	"gin-basic-framework/config"
	nsq2 "gin-basic-framework/nsq"
	"github.com/nsqio/go-nsq"
	"github.com/robfig/cron/v3"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GLOBAL_VP                *viper.Viper
	GLOBAL_CONFIG            config.Config
	GLOBAL_LOGGER            *zap.Logger
	GLOBAL_DB                *gorm.DB
	GLOBAL_CRON              *cron.Cron
	GLOBAL_NSQ_CONFIG        *nsq.Config
	GLOBAL_NSQ_PRODUCER      *nsq.Producer
	GLOBAL_NSQ_CONSUMER_POOL *nsq2.ConsumerPool
	BlackCache               local_cache.Cache
	lock                     sync.RWMutex
)

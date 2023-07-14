package initialize

import (
	"gin-basic-framework/global"
	nsq2 "gin-basic-framework/nsq"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

func NsqConfig() *nsq.Config {
	config := nsq.NewConfig()
	return config
}

func Producer() *nsq.Producer {
	config := nsq.NewConfig()
	global.GLOBAL_NSQ_CONFIG = config
	producer, err := nsq.NewProducer(global.GLOBAL_CONFIG.NSQ.AddressProvider(), config)
	if err != nil {
		global.GLOBAL_LOGGER.Fatal("Error creating NSQ producer:", zap.Error(err))
	}
	return producer
}

func ConsumerPool() *nsq2.ConsumerPool {
	pool := &nsq2.ConsumerPool{
		Consumers: make([]*nsq2.Consumer, 0),
	}
	return pool
}

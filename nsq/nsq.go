package nsq

import (
	"gin-basic-framework/config"
	"github.com/nsqio/go-nsq"
	"sync"
)

type Consumer struct {
	Topic       string
	Channel     string
	Handler     nsq.Handler
	ConsumerObj *nsq.Consumer
}

type ConsumerPool struct {
	sync.Mutex
	Consumers []*Consumer
}

func (cp *ConsumerPool) AddConsumer(topic, channel string, handler nsq.Handler, nsqConfig *config.NSQ) error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}

	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(nsqConfig.AddressProvider())
	if err != nil {
		return err
	}

	c := &Consumer{
		Topic:       topic,
		Channel:     channel,
		Handler:     handler,
		ConsumerObj: consumer,
	}

	cp.Lock()
	cp.Consumers = append(cp.Consumers, c)
	cp.Unlock()

	return nil
}

func (cp *ConsumerPool) StopAllConsumers() {
	cp.Lock()
	defer cp.Unlock()

	for _, consumer := range cp.Consumers {
		consumer.ConsumerObj.Stop()
	}
}

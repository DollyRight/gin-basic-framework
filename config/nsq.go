package config

import "strconv"

type NSQ struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (nsq *NSQ) AddressProvider() string {
	return nsq.Host + ":" + strconv.Itoa(nsq.Port)
}

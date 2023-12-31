package config

import (
	"strconv"
)

type Server struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (s *Server) AddressProvider() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

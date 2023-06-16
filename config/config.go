package config

type Config struct {
	Setting Setting `mapstructure:"setting" json:"setting" yaml:"setting"`
	Server  Server  `mapstructure:"server" json:"server" yaml:"server"`
	DB      DB      `mapstructure:"db" json:"db" yaml:"db"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

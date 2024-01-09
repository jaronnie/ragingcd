package config

var Mapping Config

type Config struct {
	DB DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Type     string
	Address  string
	Username string
	Password string
	Database string
}

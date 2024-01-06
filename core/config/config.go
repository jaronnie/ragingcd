package config

var GlobalConfig Config

type Config struct {
	DBConfig DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Type     string
	Address  string
	Username string
	Password string
	Database string
}

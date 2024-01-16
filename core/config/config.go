package config

var Mapping Config

type Config struct {
	Server `mapstructure:"server"`
}

type Server struct {
	Mode string
	DB   DBConfig `mapstructure:"db"`
	Dev  DevConfig
}

type DBConfig struct {
	Type     string
	Address  string
	Username string
	Password string
	Database string
}

type DevConfig struct {
	AuthUrl       string `mapstructure:"auth_url"`
	Authorization string
}

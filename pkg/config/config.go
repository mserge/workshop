package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	Consul  ConsulConfig
	Storage StorageConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type ConsulConfig struct {
	Hostport    string
	Ttl         string
	ServiceName string
}

type StorageConfig struct {
	Host     string
	Port     int
	Keyspace string
}

func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

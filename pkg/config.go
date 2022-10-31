package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort    string
	MongoConfig MongoConfig
}

type MongoConfig struct {
	Uri          string
	DatabaseName string
}

func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath("./configs")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &Config{
		HttpPort: v.GetString("http_port"),
		MongoConfig: MongoConfig{
			Uri:          v.GetString("mongo_uri"),
			DatabaseName: v.GetString("mongo_database_name"),
		},
	}
}

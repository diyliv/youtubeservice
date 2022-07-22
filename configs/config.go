package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port string
}

func ReadConfig() *Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error while reading config: %v\n", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Printf("Error while unmarshalling: %v\n", err)
	}

	return &cfg
}

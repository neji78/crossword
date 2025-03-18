package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type JWTConfig struct {
	Secret string
}

type DatabaseConfig struct {
	Type string
	Path string
}

func LoadConfig() *Config {
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into config struct: %v", err)
	}

	return &config
}
package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file", err)
	}

	fmt.Println("Config loaded successfully")
}

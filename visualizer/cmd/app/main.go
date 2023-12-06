package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Error while initializating configs: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

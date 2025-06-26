package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing without it...")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("The config file has been changed:", e.Name)
	})

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
}

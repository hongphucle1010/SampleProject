package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing without it...")
	}

	// Load config from config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Expand environment variables in config
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	expandEnvVars(viper.AllSettings())

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("The config file has been changed:", e.Name)
	})

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
}

// Recursive expansion of ${VAR} strings from env
func expandEnvVars(settings map[string]any) {
	for key, value := range settings {
		switch v := value.(type) {
		case string:
			if strings.Contains(v, "${") {
				settings[key] = os.ExpandEnv(v)
			}
		case map[string]any:
			expandEnvVars(v)
		}
	}
}

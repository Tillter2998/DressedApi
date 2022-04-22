package Config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB_NAME       string
	DB_COLLECTION string
	DB_USERNAME   string
	DB_PASSWORD   string
}

func NewConfig() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("database.name", "test_db")

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &config
}

package Config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB_NAME           string
	DB_COLLECTION     string
	DB_DOCUMENTS      string
	DB_USERNAME       string
	DB_PASSWORD       string
	ENVIRONMENT       string
	CERTFILE_LOCATION string
	KEYFILE_LOCATION  string
}

// Setup config
func NewConfig() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &config
}

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	URL string
}

func GetEnvVars() Config {
	if _, err := os.Stat(".env"); err == nil {
		// Initialize Viper from .env file
		viper.SetConfigFile(".env") // Specify the name of your .env file

		// Read the .env file
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading .env file: %s\n", err)
			os.Exit(1)
		}
	}

	// Enable reading environment variables
	viper.AutomaticEnv()

	// get URL from Viper
	url := viper.GetString("URL")
	if url == "" {
		fmt.Println("URL to parse must be provided")
		os.Exit(1)
	}

	var conf Config
	viper.Unmarshal(&conf)

	return conf
}

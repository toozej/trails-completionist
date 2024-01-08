package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	InputFilename     string
	ChecklistFilename string
	HTMLFilename      string
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

	// get vars from Viper
	inputFilename := viper.GetString("INPUT_FILENAME")
	if inputFilename == "" {
		fmt.Println("Input filename must be provided")
		os.Exit(1)
	}
	checklistFilename := viper.GetString("CHECKLIST_FILENAME")
	if checklistFilename == "" {
		fmt.Println("Checklist filename must be provided")
		os.Exit(1)
	}
	htmlFilename := viper.GetString("HTML_FILENAME")
	if htmlFilename == "" {
		fmt.Println("html filename must be provided")
		os.Exit(1)
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("Error unmarshalling Viper conf: %s\n", err)
		os.Exit(1)
	}

	return conf
}

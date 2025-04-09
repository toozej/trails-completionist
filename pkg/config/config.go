package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	OSMRegionFile string `mapstructure:"osm_region_file"`
	TrackFiles    string `mapstructure:"track_files"`
	InputFile     string `mapstructure:"input_file"`
	ChecklistFile string `mapstructure:"checklist_file"`
	HTMLFile      string `mapstructure:"html_file"`
	Serve         bool   `mapstructure:"serve"`
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

	// Setup conf struct with items from environment variables
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("Error unmarshalling Viper conf: %s\n", err)
		os.Exit(1)
	}

	return conf
}

// ConfigFromViper returns a Config struct populated from Viper (flags/env)
func ConfigFromViper() Config {
	return Config{
		TrackFiles:    viper.GetString("trackFiles"),
		OSMRegionFile: viper.GetString("osmRegionFile"),
		InputFile:     viper.GetString("inputFile"),
		ChecklistFile: viper.GetString("checklistFile"),
		HTMLFile:      viper.GetString("htmlFile"),
		Serve:         viper.GetBool("serve"),
	}
}

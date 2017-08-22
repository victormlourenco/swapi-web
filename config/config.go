package config

import (
	"github.com/spf13/viper"
)

func init() {
	setDefaultValues()
	readFile()
}

func setDefaultValues() {
	viper.SetDefault("ServerAddress", "localhost:8080")
	viper.SetDefault("RemoteUrl", "https://swapi.co/api/")
}

func readFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.ReadInConfig()          // Find and read the config file
}

// GetString : Return config value as string
func GetString(name string) string {
	return viper.GetString(name)
}

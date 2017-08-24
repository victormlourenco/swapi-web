package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigDefaultValues(t *testing.T) {
	setDefaultValues()
	serverAddress := viper.GetString("ServerAddress")
	if serverAddress != "0.0.0.0:6060" {
		t.Fatalf("expected a %s got %s", "0.0.0.0:6060", serverAddress)
	}
	remoteURL := viper.GetString("RemoteUrl")
	if remoteURL != "https://swapi.co/api" {
		t.Fatalf("expected a %s got %s", "https://swapi.co/api", remoteURL)
	}
}

func TestConfigFileRead(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../.")
	err := readFile()
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	serverAddress := viper.GetString("ServerAddress")
	if serverAddress == "" {
		t.Fatalf("expected a ServerAddress got null")
	}
	remoteURL := viper.GetString("RemoteUrl")
	if remoteURL != "https://swapi.co/api" {
		t.Fatalf("expected a remoteURL got null")
	}
}

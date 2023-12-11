package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config represents the configuration for the Jenkins CLI application.
// The configuration includes the url, port, user, token, and job name.
type Config struct {
	URL   string
	PORT  string
	USER  string
	TOKEN string
	JOB   string
}

// InitConfig initializes the configuration for the Jenkins CLI application.
func InitConfig() *Config {
	if err := viper.BindEnv("server", "SERVER"); err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}
	if err := viper.BindEnv("port", "PORT"); err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}
	if err := viper.BindEnv("user", "USER"); err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}
	if err := viper.BindEnv("token", "TOKEN"); err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}
	if err := viper.BindEnv("job", "JOB"); err != nil {
		log.Fatalf("Error binding environment variable: %v", err)
	}

	return &Config{
		URL:   viper.GetString("server"),
		PORT:  viper.GetString("port"),
		USER:  viper.GetString("user"),
		TOKEN: viper.GetString("token"),
		JOB:   viper.GetString("job"),
	}
}

// UpdateConfigFromArgs updates the configuration from command-line arguments.
func UpdateConfigFromArgs(args []string) *Config {
	cfg := InitConfig()
	for i, arg := range args {
		switch i {
		case 0:
			cfg.URL = arg
		case 1:
			cfg.PORT = arg
		case 2:
			cfg.USER = arg
		case 3:
			cfg.TOKEN = arg
		case 4:
			cfg.JOB = arg
		}
	}
	return cfg
}

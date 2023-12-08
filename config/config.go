package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	URL   string
	PORT  string
	USER  string
	TOKEN string
	JOB   string
}

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

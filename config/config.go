package config

import "github.com/spf13/viper"

type Config struct {
	URL   string
	PORT  string
	USER  string
	TOKEN string
	JOB   string
}

func InitConfig() *Config {
	viper.BindEnv("server", "SERVER")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("user", "USER")
	viper.BindEnv("token", "TOKEN")
	viper.BindEnv("job", "JOB")

	return &Config{
		URL:   viper.GetString("server"),
		PORT:  viper.GetString("port"),
		USER:  viper.GetString("user"),
		TOKEN: viper.GetString("token"),
		JOB:   viper.GetString("job"),
	}
}

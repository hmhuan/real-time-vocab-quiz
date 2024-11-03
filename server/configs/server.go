package configs

import (
	"github.com/spf13/viper"
)

func GetENV() string {
	env := viper.GetString("ENV")
	if env == "" {
		return "LOCAL"
	}
	return env
}

func GetHost() string {
	host := viper.GetString("SERVER_HOST")
	if host == "" {
		return "localhost"
	}
	return host
}

func GetPort() string {
	port := viper.GetString("SERVER_PORT")
	if port == "" {
		return "8080"
	}
	return port
}

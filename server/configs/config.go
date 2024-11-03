package configs

import (
	"github.com/spf13/viper"
)

func LoadConfigs() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	return nil
}
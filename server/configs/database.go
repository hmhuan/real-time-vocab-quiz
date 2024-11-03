package configs

import "github.com/spf13/viper"

func GetDBHost() string {
	return viper.GetString("DB_HOST")
}

func GetDBPort() string {
	return viper.GetString("DB_PORT")
}

func GetDBSchema() string {
	return viper.GetString("DB_SCHEMA")
}

func GetDBName() string {
	return viper.GetString("DB_NAME")

}

func GetDBUser() string {
	return viper.GetString("DB_USER")
}

func GetDBPassword() string {
	return viper.GetString("DB_PASSWORD")
}
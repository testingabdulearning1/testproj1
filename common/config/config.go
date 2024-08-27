package config

import "github.com/spf13/viper"

var (
	PORT        = "3000"
	ENVIRONMENT = ""
	SERVER_TYPE = ""
)

func LoadConfig() {
	PORT = viper.GetString("PORT")
	ENVIRONMENT = viper.GetString("ENVIRONMENT")
	SERVER_TYPE = viper.GetString("SERVER_TYPE")
}

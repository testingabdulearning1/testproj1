package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	PORT        = "3000"
	ENVIRONMENT = ""
	SERVER_TYPE = ""
)

type Config struct {
	DB string
}

func LoadConfig() {
	PORT = viper.GetString("PORT")
	ENVIRONMENT = viper.GetString("ENVIRONMENT")
	SERVER_TYPE = viper.GetString("SERVER_TYPE")

	log.Println("Config loding is success")
}

var AccountConfig struct {
	Port string `mapstructure:"ACCOUNTS_PORT"`
}

func InitAccountConfig() {
	AccountConfig.Port = os.Getenv("ACCOUNTS_PORT")
}

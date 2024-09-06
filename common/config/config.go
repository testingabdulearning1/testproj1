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

type DB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

func LoadConfig() {
	PORT = viper.GetString("PORT")
	ENVIRONMENT = viper.GetString("ENVIRONMENT")
	SERVER_TYPE = viper.GetString("SERVER_TYPE")

	log.Println("Config loading is success")
}

var AccountSvc struct {
	HttpPort string `mapstructure:"ACCOUNTS_PORT"`
	DB `mapstructure:",squash"`
}

func InitAccountConfig() {
	AccountSvc.Port = os.Getenv("ACCOUNTS_PORT")
}

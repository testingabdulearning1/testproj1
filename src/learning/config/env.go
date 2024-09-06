package config

import (
	"log"

	"github.com/spf13/viper"
	cfg "school-management-app/common/config"
)

var Env struct {
	Port string `mapstructure:"PORT"`
}

var PostgresConn cfg.PostgresConn
var PostgresPublicDbName string


func init() {
	loadConfig()
	initPublicDB()
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./internal/config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error occured while reading env variables, error:", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	err = viper.Unmarshal(&PostgresConn)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	//get value for public db
	PostgresPublicDbName = viper.GetString("POSTGRES_PUBLIC_DB_NAME")
	if PostgresPublicDbName == "" {
		log.Fatalln("error occured while writing env values onto variables, error: POSTGRES_PUBLIC_DB_NAME is empty")
	}
}

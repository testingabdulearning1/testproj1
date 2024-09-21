package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	loadConfig()
	setPublicDB()
}

var Env struct {
	Port string `mapstructure:"PORT"`
}

func loadConfig() {
	viper.SetConfigName("common")
	viper.AddConfigPath("./config")
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

	err = viper.Unmarshal(&PostgresConnection)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	PostgresPublicDbName = viper.GetString("PUBLIC_DB_NAME")
}

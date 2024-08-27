package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ImportEnv() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("SERVER_TYPE", "user")
	viper.SetDefault("ENVIRONMENT", "development")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found ignoring error
		} else {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}

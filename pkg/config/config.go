package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

var Env struct {
	Port         string `mapstructure:"PORT"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var InitialData struct {
	SuperAdminUsername string //get it directly from env variables if provided (not from .env file)
	SuperAdminPassword string //get it directly from env variables if provided (not from .env file)
}

var PostgresConn struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

func loadConfig() {
	viper.AutomaticEnv()

	if viper.GetString("DOCKER_ENV") != "true" {

		viper.SetConfigName(".env")
		viper.AddConfigPath(".")
		viper.SetConfigType("env")
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
	} else {
		Env.Port = viper.GetString("PORT")
		Env.JwtSecretKey = viper.GetString("JWT_SECRET_KEY")

		PostgresConn.DbHost = viper.GetString("DB_HOST")
		PostgresConn.DbUser = viper.GetString("DB_USER")
		PostgresConn.DbPassword = viper.GetString("DB_PASSWORD")
		PostgresConn.DbPort = viper.GetString("DB_PORT")
		PostgresConn.DbName = viper.GetString("DB_NAME")
	}

	//directly accessing the values from env variables
	InitialData.SuperAdminUsername = viper.GetString("SUPER_ADMIN_USERNAME")
	InitialData.SuperAdminPassword = viper.GetString("SUPER_ADMIN_PASSWORD")

	fmt.Println("Envirnment variables loaded successfully")
}

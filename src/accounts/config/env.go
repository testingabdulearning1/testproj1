package config

import (
	"log"

	cfg "school-management-app/common/config"

	"github.com/spf13/viper"
)

var Env struct {
	Port string `mapstructure:"PORT"`
}

var JWT struct {
	ExpTimeInMinutes int `mapstructure:"JwtTokenExpiryInMinutes"`
}

var Twilio struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
}

var PostgresConn cfg.PostgresConn
var PostgresPublicDbName string

func init() {
	loadConfig()
}

func loadConfig() {
	viper.SetConfigName("accounts")
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

	err = viper.Unmarshal(&Twilio)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	err = viper.Unmarshal(&JWT)
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

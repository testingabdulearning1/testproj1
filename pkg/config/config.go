package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

var Env struct {
	Port string `mapstructure:"PORT"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var PostgresConn struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

func loadConfig() {
	// Specify the name of the config file (without the extension)
	viper.SetConfigName(".env") // Look for a file named '.env'
	
	// Specify the path where viper should look for the config file
	viper.AddConfigPath(".")    // Look in the current directory
	
	// Set the file type for Viper to expect
	viper.SetConfigType("env")  // Set the config file type to ".env"
	
	// Automatically override any variable from the env if it exists
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

}

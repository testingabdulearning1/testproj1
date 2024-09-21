package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PublicDB *gorm.DB

var (
	PostgresConnection   PostgresConn
	PostgresPublicDbName string
)

func setPublicDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		PostgresConnection.DbHost,
		PostgresConnection.DbUser,
		PostgresConnection.DbPassword,
		PostgresPublicDbName,
		PostgresConnection.DbPort,
	)
	PublicDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Couldn't connect to DB. Error:", err)
	}
}

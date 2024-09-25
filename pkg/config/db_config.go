package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDB connects to the database and returns the connection.
// If the database does not exist, it creates it and returns the connection.
func ConnectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		PostgresConn.DbHost,
		PostgresConn.DbUser,
		PostgresConn.DbPassword,
		PostgresConn.DbName,
		PostgresConn.DbPort,
	)
	var err error

	// Create a connection string without specifying a database to connect to the PostgreSQL server
	serverDSN := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		PostgresConn.DbHost,
		PostgresConn.DbUser,
		PostgresConn.DbPassword,
		PostgresConn.DbPort,
	)

	// Connect to the PostgreSQL server
	serverDB, err := gorm.Open(postgres.Open(serverDSN), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to PostgreSQL server. Error:", err)
		return nil, err
	}

	// Check if the database exists
	var exists bool
	checkDBQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", PostgresConn.DbName)
	if err := serverDB.Raw(checkDBQuery).Scan(&exists).Error; err != nil {
		log.Printf("Couldn't check if database (named %s) exists. Error:%v\n",PostgresConn.DbName, err)
		log.Println("More info: checkDBQuery:", checkDBQuery)
		return nil, err
	}

	// If the database does not exist, create it
	if !exists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", PostgresConn.DbName)
		if err := serverDB.Exec(createDBQuery).Error; err != nil {
			log.Println("Couldn't create database. Error:", err)
			return nil, err
		}
		log.Println("Database created successfully:", PostgresConn.DbName)
	} else {
		log.Println("Database already exists:", PostgresConn.DbName)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to DB. Error:", err)
		log.Println("More info: dsn:", dsn)
		return nil, err
	}
	return db, nil
}

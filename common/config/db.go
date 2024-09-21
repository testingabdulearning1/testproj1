package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConn struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
}
type Postgresdb struct {
	PostgresConn `mapstructure:",squash"`
	DbName       string `mapstructure:"DB_NAME"`
}

// ConnectToDB connects to the database and returns the connection.
// If the database does not exist, it creates it and returns the connection.
func ConnectToDB(dbConfig Postgresdb) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
	)
	var err error

	// Create a connection string without specifying a database to connect to the PostgreSQL server
	serverDSN := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbPort,
	)

	// Connect to the PostgreSQL server
	serverDB, err := gorm.Open(postgres.Open(serverDSN), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to PostgreSQL server. Error:", err)
		return nil, err
	}

	// Check if the database exists
	var exists bool
	checkDBQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbConfig.DbName)
	if err := serverDB.Raw(checkDBQuery).Scan(&exists).Error; err != nil {
		log.Println("Couldn't check if database exists. Error:", err)
		return nil, err
	}

	// If the database does not exist, create it
	if !exists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbConfig.DbName)
		if err := serverDB.Exec(createDBQuery).Error; err != nil {
			log.Println("Couldn't create database. Error:", err)
			return nil, err
		}
		log.Println("Database created successfully:", dbConfig.DbName)
	} else {
		log.Println("Database already exists:", dbConfig.DbName)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to DB. Error:", err)
		return nil, err
	}
	return db, nil
}

// ConnectAndGetDB connects to the already existing database and returns the connection.
// If the database does not exist, it returns an error.
func ConnectAndGetDB(dbConfig Postgresdb) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// CreateNewDB creates a new database and returns the connection.
// Returns an error if the database already exists.
func CreateNewDB(dbConfig Postgresdb) (*gorm.DB, error) {
	var err error

	// Create a connection string without specifying a database to connect to the PostgreSQL server
	serverDSN := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbPort,
	)

	// Connect to the PostgreSQL server
	serverDB, err := gorm.Open(postgres.Open(serverDSN), &gorm.Config{})
	if err != nil {
		log.Println("Couldn't connect to PostgreSQL server. Error:", err)
		return nil, err
	}

	// Check if the database exists
	var exists bool
	checkDBQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", dbConfig.DbName)
	if err := serverDB.Raw(checkDBQuery).Scan(&exists).Error; err != nil {
		log.Println("Couldn't check if database exists. Error:", err)
		return nil, err
	}

	// If the database does not exist, create it
	if !exists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbConfig.DbName)
		if err := serverDB.Exec(createDBQuery).Error; err != nil {
			log.Println("Couldn't create database. Error:", err)
			return nil, err
		}
		log.Println("Database created successfully:", dbConfig.DbName)
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			dbConfig.DbHost,
			dbConfig.DbUser,
			dbConfig.DbPassword,
			dbConfig.DbName,
			dbConfig.DbPort,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Couldn't connect to DB. Error:", err)
			return nil, err
		}
		return db, nil
	} else {
		log.Println("Database already exists:", dbConfig.DbName)
		return nil, fmt.Errorf("database already exists")
	}

}

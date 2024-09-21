package config

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	DbSuffix = ""  //empty (instead of service name in microservice architecture)
)

var SchoolDBs map[string]*gorm.DB = make(map[string]*gorm.DB)

// GetSchoolDB gets the dbconnection based on the prefix code
func GetSchoolDB(schoolPrefix string) (*gorm.DB, error) {
	if db, ok := SchoolDBs[schoolPrefix]; ok {
		return db, nil
	}
	db, err := ConnectAndGetDB(Postgresdb{
		PostgresConn: PostgresConnection,
		DbName:       schoolPrefix + DbSuffix,
	})
	if err != nil {
		fmt.Println("Couldn't connect to the database. Error:", err)
		return nil, err
	}
	SchoolDBs[schoolPrefix] = db
	return db, nil
}

package db

import (
	"fmt"
	"log"

	cfg "school-management-app/common/config"
	"school-management-app/src/accounts/config"
	"school-management-app/src/accounts/domain/entities/models"

	"gorm.io/gorm"
)

var PublicDB *gorm.DB = initPublicDB()

func initPublicDB() *gorm.DB { //is called in init function in env.go
	db, err := cfg.ConnectToDB(cfg.Postgresdb{
		PostgresConn: config.PostgresConn,
		DbName:       config.PostgresPublicDbName,
	})

	if err != nil {
		log.Fatal("Couldn't connect to the database. Error:", err)
	}
	migratePublicTables(db)
	return db
}

func migratePublicTables(db *gorm.DB) {
	if err := db.AutoMigrate(&models.School{}); err != nil {
		log.Fatal("Couldn't migrate models.School. Error:", err)
	}

	if err := db.AutoMigrate(&models.SuperAdmin{}); err != nil {
		log.Fatal("Couldn't migrate models.SuperAdmin. Error:", err)
	}

	fmt.Println("Migrated tables successfully")
}

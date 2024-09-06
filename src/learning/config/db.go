package config

import (
	"fmt"
	"log"
	cfg "school-management-app/common/config"
	"school-management-app/src/learning/domain/entities/models"

	"gorm.io/gorm"
)

var PublicDB *gorm.DB

func initPublicDB() { //is called in init function in env.go
	db, err := cfg.ConnectToDB(cfg.Postgresdb{
		PostgresConn: PostgresConn,
		DbName:       PostgresPublicDbName,
	})

	if err != nil {
		log.Fatal("Couldn't connect to the database. Error:", err)
	}
	migratePublicTables(db)
	PublicDB = db
}

func migratePublicTables(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Subject{}); err != nil {
		log.Fatalln("Couldn't migrate models.Subject. Error:", err)
		return
	}

	if err := db.AutoMigrate(&models.Standards{}); err != nil {
		log.Fatalln("Couldn't migrate models.Standards. Error:", err)
		return
	}

	if err := db.AutoMigrate(&models.Syllabus{}); err != nil {
		log.Fatalln("Couldn't migrate models.Syllabus. Error:", err)
		return
	}

	if err := db.AutoMigrate(&models.Books{}); err != nil {
		log.Fatalln("Couldn't migrate models.Books. Error:", err)
		return 
	}

	fmt.Println("Migrated tables successfully")
}

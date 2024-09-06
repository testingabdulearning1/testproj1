package config

import (
	"fmt"
	cfg "school-management-app/common/config"
	"school-management-app/src/exam/domain/entities/models"

	"gorm.io/gorm"
)

const (
	DbSuffix = "_exam"
)

var SchoolDBs map[string]*gorm.DB = make(map[string]*gorm.DB)

func GetSchoolDB(schoolPrefix string) (*gorm.DB, error) {
	if db, ok := SchoolDBs[schoolPrefix]; ok {
		return db, nil
	}
	db, err := cfg.ConnectAndGetDB(cfg.Postgresdb{
		PostgresConn: PostgresConn,
		DbName:       schoolPrefix + DbSuffix,
	})
	if err != nil {
		fmt.Println("Couldn't connect to the database. Error:", err)
		return nil, err
	}
	SchoolDBs[schoolPrefix] = db
	return db, nil
}
func CreateNewSchool(schoolPrefix string) error {
	db, err := cfg.CreateNewDB(cfg.Postgresdb{
		PostgresConn: PostgresConn,
		DbName:       schoolPrefix + DbSuffix,
	})
	if err != nil {
		fmt.Println("Couldn't connect to the database. Error:", err)
		return err
	}

	if err := migrateExamTables(db); err != nil {
		fmt.Println("Couldn't migrate accounts tables. Error:", err)
		return err
	}

	SchoolDBs[schoolPrefix] = db
	return nil
}

// This function is to be called when a new school is created, to create the exam tables for the school
func migrateExamTables(schoolDb *gorm.DB) error {
	if err := schoolDb.AutoMigrate(&models.TermExam{}); err != nil {
		fmt.Println("Couldn't migrate models.TermExam. Error:", err)
		return err
	}

	fmt.Println("Migrated tables successfully")
	return nil
}

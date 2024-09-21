package db

import (
	"fmt"
	cfg "school-management-app/common/config"
	"school-management-app/src/attendance/config"
	"school-management-app/src/attendance/domain/entities/models"

	"gorm.io/gorm"
)

const (
	DbSuffix = "_attendance"
)

var SchoolDBs map[string]*gorm.DB = make(map[string]*gorm.DB)

func GetSchoolDB(schoolPrefix string) (*gorm.DB, error) {
	if db, ok := SchoolDBs[schoolPrefix]; ok {
		return db, nil
	}
	db, err := cfg.ConnectAndGetDB(cfg.Postgresdb{
		PostgresConn: config.PostgresConn,
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
		PostgresConn: config.PostgresConn,
		DbName:       schoolPrefix + DbSuffix,
	})
	if err != nil {
		fmt.Println("Couldn't connect to the database. Error:", err)
		return err
	}

	if err := migrateAttendanceTables(db); err != nil {
		fmt.Println("Couldn't migrate accounts tables. Error:", err)
		return err
	}

	SchoolDBs[schoolPrefix] = db
	return nil
}

// This function is to be called when a new school is created, to create the attendance tables for the school
func migrateAttendanceTables(schoolDb *gorm.DB) error {
	if err := schoolDb.AutoMigrate(&models.StudentAttendanceEdits{}); err != nil {
		fmt.Println("Couldn't migrate models.StudentAttendanceEdits. Error:", err)
		return err
	}

	fmt.Println("Migrated tables successfully")
	return nil
}

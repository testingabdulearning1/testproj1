package db

import (
	"fmt"
	"log"
	"school-management-app/pkg/config"
	"school-management-app/pkg/domain/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	superAdminUsername = "superadmin"
	superAdminPassword = "adminpwsuper"
)

var PublicDB *gorm.DB = initPublicDB()

func initPublicDB() *gorm.DB { //is called in init function in env.go
	db, err := config.ConnectToDB()

	if err != nil {
		log.Fatal("Couldn't connect to the database. Error:", err)
	}
	migratePublicTables(db)
	initiateSuperAdmin(db)
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

func initiateSuperAdmin(db *gorm.DB) {
	//check if super admin exists
	var superAdmin models.SuperAdmin
	err := db.First(&superAdmin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			encryptedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(superAdminPassword), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Couldn't encrypt super admin password. Error:", err)
			}

			//create super admin
			superAdmin = models.SuperAdmin{
				Username: superAdminUsername,
				HashedPassword: string(encryptedPasswordByte),
			}
			err = db.Create(&superAdmin).Error
			if err != nil {
				log.Fatal("Couldn't create super admin. Error:", err)
			}
			fmt.Println("Super admin created successfully")
		} else {
			log.Fatal("Couldn't get super admin. Error:", err)
		}
	}else {
		fmt.Println("Super admin already exists")
	}
}

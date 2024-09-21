package repositories

import (
	interfacess "school-management-app/pkg/repository/interface"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfacess.Repository {
	return Repo{db: db}
}

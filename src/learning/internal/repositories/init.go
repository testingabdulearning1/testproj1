package repositories

import (
	"school-management-app/src/learning/domain/repositories"

	"gorm.io/gorm"
)

type CurriculumRepo struct {
	PublicDB *gorm.DB
}

func NewCurriculumRepo(publicDb *gorm.DB) repositories.CurriculumRepo { 
	return &CurriculumRepo{PublicDB: publicDb}
}

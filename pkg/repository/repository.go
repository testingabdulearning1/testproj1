package repositories

import (
	"context"
	"school-management-app/pkg/domain/models"
	interfacess "school-management-app/pkg/repository/interface"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfacess.Repository {
	return Repo{db: db}
}

// GetSuperAdminPassword returns the hashed password of the super admin
func (r Repo) GetSuperAdminPassword(ctx context.Context, username string) (bool, string, error) {
	var hashedPassword string
	err := r.db.Model(&models.SuperAdmin{}).Where("username = ?", username).Select("hashed_password").Scan(&hashedPassword).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, "", nil
		}
		return false, "", err
	}
	return true, hashedPassword, nil
}

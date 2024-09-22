package interfaces

import (
	"context"
	"school-management-app/pkg/domain/models"
)

type Repository interface {
	GetSuperAdminPassword(ctx context.Context, username string) (*models.SuperAdmin,bool, error)
}

package interfaces

import "context"

type Repository interface {
	GetSuperAdminPassword(ctx context.Context, username string) (bool,string, error)
}

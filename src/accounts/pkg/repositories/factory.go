package repositories

import "school-management-app/src/accounts/domain/repositories"

type UserRepo struct {
	// db *sql.DB
	db interface{} //change db type to interface{} from config
}

func NewUserRepo(db interface{}) repositories.AuthRepository { //change interface{} to db type from config
	return &UserRepo{db: db}
}

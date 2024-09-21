package usecase

import (
	"context"
	"net/http"
	"school-management-app/pkg/domain/request"
	"school-management-app/pkg/domain/respcode"
	"school-management-app/pkg/domain/response"
	interfaces "school-management-app/pkg/repository/interface"
	usecases "school-management-app/pkg/usecase/interface"

	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo interfaces.Repository
}

func NewUsecase(repo interfaces.Repository) usecases.Usecases {
	return Usecase{repo: repo}
}

func (u Usecase) SuperAdminSignin(ctx context.Context, req *request.SuperAdminSigninReq) response.Response {
	//get super admin by username
	usernameExists, hashedPassword, err := u.repo.GetSuperAdminPassword(ctx, req.Username)
	if err != nil {
		return response.CreateError(http.StatusInternalServerError, respcode.InternalServerError, err)
	}
	if !usernameExists {
		return response.CreateError(http.StatusUnauthorized, respcode.InvalidUsername, nil)
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		return response.CreateError(http.StatusUnauthorized, respcode.InvalidPassword, nil)
	}

	return response.CreateSuccess(http.StatusOK, respcode.Success, nil)
}
package usecase

import (
	"context"
	"fmt"
	"net/http"
	"school-management-app/pkg/domain/constants"
	"school-management-app/pkg/domain/request"
	"school-management-app/pkg/domain/respcode"
	"school-management-app/pkg/domain/response"
	interfaces "school-management-app/pkg/repository/interface"
	usecases "school-management-app/pkg/usecase/interface"
	jwttoken "school-management-app/pkg/utils/jwt-token"

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
	admin, recordExists, err := u.repo.GetSuperAdminPassword(ctx, req.Username)
	if err != nil {
		return response.CreateError(http.StatusInternalServerError, respcode.InternalServerError, err)
	}
	if !recordExists {
		return response.CreateError(http.StatusUnauthorized, respcode.UsernameNotRegd, nil)
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.HashedPassword), []byte(req.Password)); err != nil {
		return response.CreateError(http.StatusUnauthorized, respcode.PasswordMismatch, nil)
	}

	//generate token
	token, err := jwttoken.GenerateToken(admin.ID, constants.RoleSuperAdmin, nil)
	if err != nil {
		return response.CreateError(http.StatusInternalServerError, respcode.InternalServerError, fmt.Errorf("error in generating token: %v", err))
	}
	
	return response.Response{
		HttpStatusCode: http.StatusOK,
		Status:         true,
		ResponseCode:   respcode.Success,
		Data: map[string]interface{}{
			"name":  admin.Name,
			"token": token,
		},
	}
}

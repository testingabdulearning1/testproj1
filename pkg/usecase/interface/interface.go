package usecases

import (
	"context"
	"school-management-app/pkg/domain/request"
	"school-management-app/pkg/domain/response"
)

type Usecases interface {
	SuperAdminSignin(ctx context.Context, req *request.SuperAdminSigninReq) response.Response
}

package usecases

import (
	"errors"
	"fmt"
	"school-management-app/common/domain/respcode"
	"school-management-app/common/domain/response"
	"school-management-app/src/learning/domain/entities/request"
)

func (u *CurriculumUseCase) Sample(req request.SampleReq) response.Response {
	str, err := u.CurriculumRepo.Ping(req)
	if err != nil {
		return response.Response{
			HttpStatusCode: 300,
			Status:         true,
			ResponseCode:   respcode.Unauthorized,
			Error:          errors.New("database error"),
			Data:           nil,
		}
	}
	fmt.Println("pong: ", str)
	// Add your logic here
	return response.Response{
		HttpStatusCode: 300,
		Status:         true,
		ResponseCode:   respcode.Unauthorized,
		Error:          errors.New("database error"),
		Data:           nil,
	}
}

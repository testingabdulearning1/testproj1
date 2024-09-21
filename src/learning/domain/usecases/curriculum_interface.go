package usecases

import (
	"school-management-app/common/domain/response"
	"school-management-app/src/learning/domain/entities/request"
)

type CurriculumUC interface {
	Sample(req request.SampleReq) response.Response
}

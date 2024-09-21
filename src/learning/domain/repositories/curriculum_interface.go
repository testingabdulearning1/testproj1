package repositories

import "school-management-app/src/learning/domain/entities/request"

type CurriculumRepo interface {
	Ping(request.SampleReq) (string,error)
}

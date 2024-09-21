package repositories

import (
	"school-management-app/src/learning/domain/entities/request"
)
func (repo *CurriculumRepo) Ping(req request.SampleReq) (string, error) {
	// schoolPrefix := "ABCD"

	return "pong", nil
}

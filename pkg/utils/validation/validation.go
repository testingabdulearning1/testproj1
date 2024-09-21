package validation

import (
	"fmt"
	"school-management-app/pkg/domain/response"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func validateRequestDetailed(req interface{}) []response.InvalidField {

	Response := []response.InvalidField{}
	errs := validate.Struct(req)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		e := response.InvalidField{
			FailedField: err.Field(),
			Tag:         err.Tag(),
			Value:       err.Value(),
		}

		message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)
		fmt.Println("validation fail message: ", message)

		Response = append(Response, e)
	}
	return Response
}

// func ValidateRequest(req interface{}) []string {
// 	Response := []string{}
// 	errs := validate.Struct(req)

// 	if errs == nil {
// 		return nil
// 	}

// 	for _, err := range errs.(validator.ValidationErrors) {
// 		e := response.InvalidField {
// 			// Error:       true,
// 			FailedField: err.Field(),
// 			Tag:         err.Tag(),
// 			Value:       err.Value(),
// 		}

// 		message := fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", e.FailedField, e.Value, e.Tag)

// 		Response = append(Response, message)
// 	}
// 	return Response
// }

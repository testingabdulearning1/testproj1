package validation

import (
	"log"
	"net/http"
	"school-management-app/pkg/domain/respcode"
	"school-management-app/pkg/domain/response"

	"github.com/gofiber/fiber/v2"
)

// BindAndValidateRequest binds and validates the request.
// Req should be a pointer to the request struct.
func BindAndValidateRequest(c *fiber.Ctx, req interface{}) (bool, error) {
	if err := c.BodyParser(req); err != nil {
		log.Println("error parsing request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.Response{
			Status:       false,
			ResponseCode: respcode.BindingError,
			Error:        err,
		})
	}

	if err := validateRequestDetailed(req); err != nil {
		log.Println("error validating request:", err)
		return false, c.Status(http.StatusBadRequest).JSON(response.ValidationErrorResponse{
			Status:       false,
			ResponseCode: respcode.ValidationError,
			Errors:       err,
		})
	}
	return true, nil
}

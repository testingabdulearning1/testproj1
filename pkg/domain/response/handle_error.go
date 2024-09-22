package response

import (
	"github.com/gofiber/fiber/v2"
)

// type ResponseI interface {
// 	WriteToJSON(c *fiber.Ctx) error
// }

type custError struct {
	Response
	Error string `json:"error"`
}

func (resp Response) WriteToJSON(c *fiber.Ctx) error {
	if resp.Error == nil {
		return c.Status(resp.HttpStatusCode).JSON(resp)
	}
	newCustError := custError{
		Response: resp,
		Error:    resp.Error.Error(),
	}
	return c.Status(resp.HttpStatusCode).JSON(newCustError)
}

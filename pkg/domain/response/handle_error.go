package response

import (
	"fmt"

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
	newCustError := custError{
		Response: resp,
	}
	if resp.Error != nil {
		newCustError.Error = resp.Error.Error()
	}
	fmt.Println("newCustError", newCustError)
	return c.Status(resp.HttpStatusCode).JSON(newCustError)
}

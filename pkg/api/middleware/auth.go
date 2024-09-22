package middleware

import (
	"fmt"
	cons "school-management-app/pkg/domain/constants"
	"school-management-app/pkg/domain/respcode"
	"school-management-app/pkg/domain/response"
	jwttoken "school-management-app/pkg/utils/jwt-token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ValidateJWT validates the jwt token in the Authorization header.
//
// If the token is valid, it sets the adminID, role and addlInfo in the context. 
// Authentication middlewares for specific roles should be used after this middleware.
func ValidateJWT(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	tokenData, err := jwttoken.GetDataFromToken(tokenString)
	if err != nil {
		if err == jwttoken.ErrTokenExpired {
			return c.Status(401).JSON(response.Response{
				Status:       false,
				ResponseCode: respcode.TokenExpired,
				Error:        err,
			})
		} else {
			return c.Status(401).JSON(response.Response{
				Status:       false,
				ResponseCode: respcode.InvalidToken,
				Error:        err,
			})
		}
	}
	c.Locals("adminID", tokenData.Id)
	c.Locals("role", tokenData.Role)
	c.Locals("addlInfo", tokenData.AddlInfo)

	return c.Next()
}

// CheckIfSuperAdmin authenticates the user as a super admin.
//
// This middleware should be used after ValidateJWT middleware.
func CheckIfSuperAdmin(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != cons.RoleSuperAdmin {
		return c.Status(403).JSON(response.Response{
			Status:       false,
			ResponseCode: respcode.Forbidden,
			Error:        fmt.Errorf("role in jwt token is not '%s',but is '%s'", cons.RoleSuperAdmin, role),
		})
	}
	return c.Next()
}

// CheckIfAdmin authenticates the user as an admin.
//
// This middleware should be used after ValidateJWT middleware.
func CheckIfStudent(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != cons.RoleStudent {
		return c.Status(403).JSON(response.Response{
			Status:       false,
			ResponseCode: respcode.Forbidden,
			Error:        fmt.Errorf("role in jwt token is not '%s',but is '%s'", cons.RoleStudent, role),
		})
	}
	return c.Next()
}

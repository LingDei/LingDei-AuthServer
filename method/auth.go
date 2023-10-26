package method

import (
	"LingDei-AuthServer/model"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// GetUserRole 解析用户token凭证的用户信息
func GetUserFromToken(c *fiber.Ctx) model.User {
	var result model.User
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	result.ID = claims["id"].(string)
	result.UserName = claims["username"].(string)
	result.Role = int(claims["role"].(float64))
	return result
}

// CheckUserRole 检查用户权限登记
func CheckUserRole(role, role_required int) error {
	if role < role_required {
		return errors.New("role denied")
	}
	return nil
}

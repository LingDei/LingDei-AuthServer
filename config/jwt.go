package config

import (
	"LingDei_AuthServer/model"

	"github.com/gofiber/fiber/v2"
)

// CustomJwtError 自定义jwt报错信息
func CustomJwtError() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err.Error() == "Missing or malformed JWT" {
			return c.Status(fiber.StatusBadRequest).JSON(model.OperationResp{Code: 400, Msg: err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(model.OperationResp{Code: 400, Msg: "Invalid or expired JWT"})
	}
}

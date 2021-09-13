package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/dto"
)

func GetUserDetails(c *fiber.Ctx) error {
	return c.JSON(dto.Success(nil))
}

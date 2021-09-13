package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
)

func InitRoutes(ctx *config.AppContext, app *fiber.App) {
	app.Get("/", GetUserDetails)
}

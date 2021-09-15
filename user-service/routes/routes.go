package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
)

func InitRoutes(ctx *config.AppContext, app *fiber.App) {
	app.Get("/", GetUserDetails)

	withAppContext := app.Group("/", func(c *fiber.Ctx) error {
		c.SetUserContext(context.WithValue(c.UserContext(), config.AppContextKey{}, ctx))
		return c.Next()
	})

	versionedRouter := withAppContext.Group(ctx.Config.Verison.String(), func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	versionedRouter.Get("/user/:userId", GetUserDetails)
	versionedRouter.Put("/user", CreateUser)

}

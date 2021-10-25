//Package routes has the routes for the api
package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
)

//InitRoutes initializes routes
func InitRoutes(ctx *config.AppContext, app *fiber.App) {
	/*
	 * We will get app context from server context
	 * Will give version to api endpoints
	 * Will initialize the routes
	 */
	//getting the app context
	withAppContext := app.Group("/", func(c *fiber.Ctx) error {
		c.SetUserContext(context.WithValue(c.UserContext(), config.AppContextKey{}, ctx))
		return c.Next()
	})

	//adding version to the api endpoints
	versionedRouter := withAppContext.Group(ctx.Config.Verison.String(), func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	//initializing routes
	versionedRouter.Get("/user", GetAllUsers)
	versionedRouter.Get("/user/:userId", GetUserDetails)
	versionedRouter.Put("/user", CreateUser)
	versionedRouter.Post("/user", UpdateUser)
	versionedRouter.Post("/login", Login)

}

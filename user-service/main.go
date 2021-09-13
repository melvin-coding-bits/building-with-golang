package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/routes"
)

func main() {
	ctx, err := config.InitContext()
	if err != nil {
		log.Fatalf("error initializing context: %v", err)
	}
	ctx.Logger.Info("context loaded")

	app := fiber.New()
	routes.InitRoutes(ctx, app)
	ctx.Logger.Info("app routes initialized")

	ctx.Logger.WithField("port", ctx.Config.Server.Port).Info("starting server")
	app.Listen(ctx.Config.Server.GetPort())
}

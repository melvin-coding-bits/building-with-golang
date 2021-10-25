package main

import (
	"log"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	/*
	 * We will initialize the context
	 * Then we will initialize the routes
	 * Then we will initialize the server
	 */
	//we initialize the context
	ctx, err := config.InitContext()
	if err != nil {
		log.Fatalf("error initializing context: %v", err)
	}
	ctx.Logger.Info("context loaded")

	//then initialize the routes
	app := fiber.New()
	routes.InitRoutes(ctx, app)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	ctx.Logger.Info("app routes initialized")

	//start the server
	ctx.Logger.WithField("port", ctx.Config.Server.Port).Info("starting server")
	app.Listen(ctx.Config.Server.GetPort())
}

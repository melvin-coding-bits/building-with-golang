package main

import (
	"log"
	"net/http"

	"github.com/melvinodsa/build-with-golang/api-gateway/config"
	"github.com/melvinodsa/build-with-golang/api-gateway/proxy"
)

func main() {
	ctx, err := config.InitContext()
	if err != nil {
		log.Fatal(err)
	}
	ctx.Logger.Info("context initialized")

	proxy.InitProxies(ctx)
	ctx.Logger.Info("proxies initialized")

	ctx.Logger.WithField("port", ctx.Config.Server.Port).Info("server started")
	log.Fatal(http.ListenAndServe(ctx.Config.Server.GetPort(), nil))
}

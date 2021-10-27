package main

import (
	"io"
	"log"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/routes"
	"github.com/opentracing/opentracing-go"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func enableTracing() (io.Closer, error) {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, err
	}
	cfg.ServiceName = "user service app"
	if cfg.Sampler == nil {
		cfg.Sampler = &jaegercfg.SamplerConfig{}
	}
	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1
	if cfg.Reporter == nil {
		cfg.Reporter = &jaegercfg.ReporterConfig{}
	}
	cfg.Reporter.LogSpans = true

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	opentracing.SetGlobalTracer(tracer)
	return closer, err
}

func main() {
	/*
	 * We will initialize the context
	 * Then we will initialize the routes
	 * Then we will initialize the server
	 */
	//we initialize the context

	closer, err := enableTracing()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("loading-config")
	ctx, err := config.InitContext()
	if err != nil {
		log.Fatalf("error initializing context: %v", err)
	}
	ctx.Logger.Info("context loaded")
	span.Finish()

	//then initialize the routes
	app := fiber.New()
	routes.InitRoutes(ctx, app)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	ctx.Logger.Info("app routes initialized")

	//start the server
	ctx.Logger.WithField("port", ctx.Config.Server.Port).Info("starting server")
	app.Listen(ctx.Config.Server.GetPort())
}

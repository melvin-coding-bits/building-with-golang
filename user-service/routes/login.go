package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/melvinodsa/build-with-golang/user-service/config"
	"github.com/melvinodsa/build-with-golang/user-service/crypto_utils"
	"github.com/melvinodsa/build-with-golang/user-service/dto"
	"github.com/melvinodsa/build-with-golang/user-service/helper/db"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var loggedInUsersCount = promauto.NewCounter(prometheus.CounterOpts{
	Name: "logged_in_users",
	Help: "No. of logged in users",
})

//Login logs the user into the system
func Login(c *fiber.Ctx) error {
	/*
	 * We will get the user context
	 * We will get the auth payload to be checked
	 * We will validate the user email and password
	 * We will generate the token and set as auth token
	 */
	tracer := opentracing.GlobalTracer()
	spanA := tracer.StartSpan("login api")
	defer spanA.Finish()
	span := tracer.StartSpan("get user context", opentracing.ChildOf(spanA.Context()))
	defer span.Finish()
	ctx, ok := c.UserContext().Value(config.AppContextKey{}).(*config.AppContext)
	if !ok {
		return c.JSON(dto.Error(errors.New("context not found"), http.StatusInternalServerError))
	}
	spanBP := tracer.StartSpan("body parsing", opentracing.ChildOf(spanA.Context()))
	defer spanBP.Finish()
	a := &dto.Auth{}
	if err := c.BodyParser(a); err != nil {
		payloadParsingError.Inc()
		ctx.Logger.WithField("parsing", "body").Error(err)
		return c.JSON(dto.Error(err, http.StatusBadRequest))
	}

	spanCP := tracer.StartSpan("checking password", opentracing.ChildOf(spanA.Context()))
	defer spanCP.Finish()
	user, err := db.CheckPassword(ctx, a.Email, a.Password)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusForbidden))
	}

	spanGT := tracer.StartSpan("generating token", opentracing.ChildOf(spanA.Context()))
	defer spanGT.Finish()
	token, err := crypto_utils.GenerateToken(32)
	if err != nil {
		return c.JSON(dto.Error(err, http.StatusInternalServerError))
	}
	c.Response().Header.Set("Authorization", "Bearer "+token)
	loggedInUsersCount.Inc()

	return c.JSON(dto.Success(dto.FromUserModel(user)))
}

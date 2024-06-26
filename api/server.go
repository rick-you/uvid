package api

import (
	"time"

	"github.com/rick-you/uvid/config"
	"github.com/rick-you/uvid/daos"
	"github.com/rick-you/uvid/dtos"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	App *echo.Echo
	Dao *daos.Dao
}

var Configs = map[string]string{}

func New(dsn string) Server {
	server := Server{
		App: echo.New(),
		Dao: daos.New(dsn),
	}
	// daos.Seed(dsn)

	server.App.Validator = &dtos.CustomValidator{Validator: validator.New()}

	server.App.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	server.App.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))
	server.App.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Minute,
	}))
	server.App.Use(middleware.Recover())
	server.App.Use(middleware.Secure())
	server.App.Use(middleware.Logger())

	bindSpanApi(server)
	bindDashApi(server)
	bindDashStatic(server)
	return server
}

func (server Server) Start() {
	server.App.Logger.Fatal(server.App.Start(":" + config.CLIConfig.Port))
}

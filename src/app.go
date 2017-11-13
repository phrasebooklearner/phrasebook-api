package src

import (
	"database/sql"

	"phrasebook-api/src/config"
	"phrasebook-api/src/database"
	"phrasebook-api/src/handler/user"
	"phrasebook-api/src/repository"
	"phrasebook-api/src/errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"phrasebook-api/src/handler"
	"phrasebook-api/src/response"
)

type app struct {
	config     config.Config
	router     *echo.Echo
	db         *sql.DB
	repository struct {
		user repository.UserRepository
	}
}

func NewApp(config config.Config, router *echo.Echo) {
	app := &app{
		config: config,
		router: router,
		db:     database.NewDBConnection(config.GetDatabaseDSN()),
	}

	app.setRepositories()
	app.setCustomErrorHandling()

	app.initHandlers()
}

func (a *app) setRepositories() {
	a.repository.user = repository.NewUserRepository(a.db)
}

func (a *app) initHandlers() {
	var routers = []handler.RoutingInitiator{
		user.NewRegistrationHandler(a.repository.user),
	}
	for _, route := range routers {
		route.InitRouting(a.router)
	}
}

func (a *app) setCustomErrorHandling() {
	a.router.HTTPErrorHandler = func(err error, c echo.Context) {
		var apiErr errors.ApiError
		if tmp, ok := err.(errors.ApiError); ok {
			apiErr = tmp
		} else if tmp, ok := err.(*echo.HTTPError); ok {
			apiErr = errors.NewHTTPError(tmp.Code)
		} else {
			apiErr = errors.NewInternalError(err)
		}

		c.JSON(apiErr.GetHTTPCode(), response.ApiError(apiErr, a.config.GetDebugEnabled()))
	}
}

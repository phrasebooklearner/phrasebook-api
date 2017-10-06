package src

import (
	"database/sql"

	"phrasebook-api/src/config"
	"phrasebook-api/src/controller"
	"phrasebook-api/src/database"
	apiError "phrasebook-api/src/error"
	"phrasebook-api/src/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type app struct {
	config config.Config
	router *echo.Echo
	db     *sql.DB
	repository   struct {
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

	app.initControllers()
}

func (a *app) setRepositories() {
	a.repository.user = repository.NewUserRepository(a.db)
}

func (a *app) initControllers() {
	controller.NewUserController(a.repository.user).InitRouting(a.router)
}

func (a *app) setCustomErrorHandling() {
	defaultHandler := a.router.HTTPErrorHandler
	a.router.HTTPErrorHandler = func(err error, c echo.Context) {
		if apiErr, ok := err.(apiError.ApiError); ok {
			c.JSON(apiErr.GetHTTPCode(), map[string]map[string]interface{}{
				"error": {
					"type":        apiErr.GetErrorType(),
					"fullMessage": apiErr.Error(),
					"data":        apiErr,
				},
			})
		} else {
			defaultHandler(err, c)
		}
	}
}

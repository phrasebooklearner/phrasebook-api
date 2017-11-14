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
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
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
	app.setAuthProviders()

	app.initHandlers()
}

// Here different data repositories are created
func (a *app) setRepositories() {
	a.repository.user = repository.NewUserRepository(a.db)
}

// All handlers are located in routers slice.
// Every handler should implement handler.RoutingInitiator interface.
// Each handler is responsible for its routing settings.
func (a *app) initHandlers() {
	var routers = []handler.RoutingInitiator{
		user.NewRegistrationHandler(a.repository.user),
	}
	for _, route := range routers {
		route.InitRouting(a.router)
	}
}

// All errors that are returned to the client should implement errors.ApiError interface.
// If an error doesn't implement that interface, then it is wrapped into errors.internalServerError,
// that implements errors.ApiError interface.
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

// A github.com/markbates/goth package is used as an auth library.
func (a *app) setAuthProviders() {
	// TODO move provider arguments to config
	goth.UseProviders(
		gplus.New("44166123467-o6brs9o43tgaek9q12lef07bk48m3jmf.apps.googleusercontent.com", "rpXpakthfjPVoFGvcf9CVCu7", "http://localhost:8080/auth/callback/google"),
	)
}
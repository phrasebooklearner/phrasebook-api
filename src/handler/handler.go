package handler

import "github.com/labstack/echo"

type RoutingInitiator interface {
	InitRouting(router *echo.Echo)
}

package main

import (
	"phrasebook-api/src"
	"phrasebook-api/src/config"

	"github.com/labstack/echo"
)

func main() {
	cfg := config.NewEnvConfig()

	router := echo.New()
	router.Debug = cfg.GetDebugEnabled()

	src.NewApp(cfg, router)

	router.Logger.Fatal(router.Start(":" + cfg.GetPort()))
}

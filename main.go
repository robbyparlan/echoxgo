package main

import (
	"echoxgo/src/router"
	"net/http"

	util "echoxgo/src/utils/helper"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(util.LoggingMiddleware)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service echoxgo")
	})

	router.AllRoutes(e)

	e.Logger.Fatal(e.Start(":5003"))
}

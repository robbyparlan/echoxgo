package main

import (
	"echoxgo/src/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service echoxgo")
	})

	router.AllRoutes(e)

	e.Logger.Fatal(e.Start(":5003"))
}

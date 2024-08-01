package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userGroup := e.Group("/user")
	userGroup.GET("/list", func(c echo.Context) error {
		return c.String(http.StatusOK, "user list")
	})
}

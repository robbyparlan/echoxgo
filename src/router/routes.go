package router

import "github.com/labstack/echo/v4"

func AllRoutes(e *echo.Echo) {
	route := e.Group("/api/v1")
	UserRoutes(route)
}

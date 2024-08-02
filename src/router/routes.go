package router

import (
	utils "echoxgo/src/utils/constants"

	"github.com/labstack/echo/v4"
)

func AllRoutes(e *echo.Echo) {
	route := e.Group(utils.ApiPrefixVersion)
	UserRoutes(route)
}

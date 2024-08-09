package router

import (
	UserController "echoxgo/src/app/user/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userGroup := e.Group("/user")
	userGroup.POST("/list", UserController.ListUser())
}

package user

import (
	"net/http"

	userdtos "echoxgo/src/app/user/dtos"
	utx "echoxgo/src/utils/helper"

	"github.com/labstack/echo/v4"
)

func ListUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		reqBody := &userdtos.UserListDto{}
		if err := ctx.Bind(&reqBody); err != nil {
			utx.Logger.Error(`--------------- Error : `, err.Error())
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		utx.Logger.Info(`-------------- payload : `, reqBody)
		return ctx.JSON(http.StatusOK, reqBody)
	}
}

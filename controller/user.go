package controller

import (
	"os"

	"github.com/labstack/echo/v4"

	"github.com/namhq1989/daily-toolkit/model"
	"github.com/namhq1989/daily-toolkit/service"
	"github.com/namhq1989/daily-toolkit/util"
)

// UserGetToken ...
func UserGetToken(c echo.Context) error {
	// Return if in release mode
	if os.Getenv("ENV") == "release" {
		return util.Response404(c, echo.Map{}, "")
	}

	var (
		userID = model.FormatStringToUUID(c.QueryParam("id"))
		ctx    = c.Request().Context()
	)

	// Return if user id is nil
	if !util.IsValidUUID(userID) {
		return util.Response404(c, echo.Map{}, "")
	}

	user := service.UserGetByID(ctx, userID)

	return util.Response200(c, echo.Map{"token": user.GenerateToken()}, "")
}

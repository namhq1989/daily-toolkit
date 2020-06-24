package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/daily-toolkit/util"
)

// ExpenseListAll ...
func ExpenseListAll(c echo.Context) error {
	// var (
	// userID = util.EchoContextGetCurrentUserID(c)
	// )
	return util.Response200(c, echo.Map{"expenses": make([]string, 0)}, "")
}

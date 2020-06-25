package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/namhq1989/daily-toolkit/model"
	"github.com/namhq1989/daily-toolkit/service"
	"github.com/namhq1989/daily-toolkit/util"
)

// ExpenseListAll ...
func ExpenseListAll(c echo.Context) error {
	var (
		userID    = util.EchoContextGetCurrentUserID(c)
		pageToken = util.PageTokenDecode(c)
		query     = model.CommonQuery{
			Limit:     20,
			Timestamp: pageToken.Timestamp,
		}
		ctx = c.Request().Context()
	)

	data, _ := service.ExpenseGetByUserID(ctx, userID, query)

	return util.Response200(c, echo.Map{"expenses": data}, "")
}

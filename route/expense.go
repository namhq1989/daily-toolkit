package route

import (
	"github.com/labstack/echo/v4"

	"github.com/namhq1989/daily-toolkit/controller"
)

func expense(e *echo.Echo) {
	routes := e.Group("/expenses")

	// List expenses
	routes.GET("", controller.ExpenseListAll, requireLogin)
}

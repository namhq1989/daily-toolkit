package route

import (
	"github.com/labstack/echo/v4"

	"github.com/namhq1989/daily-toolkit/controller"
)

func user(e *echo.Echo) {
	routes := e.Group("/users")

	// List expenses
	routes.GET("/token", controller.UserGetToken)
}

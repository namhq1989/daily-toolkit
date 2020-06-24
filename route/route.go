package route

import (
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {
	e.Use(configCORS())
	e.Use(auth)

	// Boostrap routes
	expense(e)
}

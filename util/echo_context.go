package util

import (
	"github.com/labstack/echo/v4"

	"github.com/namhq1989/daily-toolkit/config"
	"github.com/namhq1989/daily-toolkit/model"
)

// EchoContextGetCurrentUserID return current user id
func EchoContextGetCurrentUserID(c echo.Context) model.UUID {
	userID := c.Get(config.EchoContextKeyCurrentUserID).(model.UUID)
	return userID
}

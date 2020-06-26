package util

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"github.com/namhq1989/daily-toolkit/config"
	"github.com/namhq1989/daily-toolkit/model"
)

// EchoContextGetCurrentUserID return current user id
func EchoContextGetCurrentUserID(c echo.Context) model.UUID {
	userIDString := c.Get(config.EchoContextKeyCurrentUserID).(string)
	userID := uuid.FromStringOrNil(userIDString)
	return userID
}

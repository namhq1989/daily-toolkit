package util

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/namhq1989/daily-toolkit/config"
)

// EchoContextGetCurrentUserID return current user id
func EchoContextGetCurrentUserID(c echo.Context) primitive.ObjectID {
	userID := c.Get(config.EchoContextKeyCurrentUserID)

	// Convert to object id
	objectID, _ := primitive.ObjectIDFromHex(userID.(string))
	return objectID
}

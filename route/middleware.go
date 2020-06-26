package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	uuid "github.com/satori/go.uuid"

	"github.com/namhq1989/daily-toolkit/config"
	"github.com/namhq1989/daily-toolkit/util"
)

// configCORS ...
func configCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderContentLength, echo.HeaderAuthorization},
		AllowCredentials: false,
		MaxAge:           600,
	})
}

// Auth api before next to controller
func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get header
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

		// Parse token
		tokenData, _ := parseToken(authHeader, "app_secret") // FIXME: change secret key

		// Skip if cannot parse token
		if tokenData == nil {
			c.Set(config.EchoContextKeyCurrentUserID, "")
			return next(c)
		}

		data, ok := tokenData.Claims.(jwt.MapClaims)
		if ok && tokenData.Valid && data["_id"] != "" {
			c.Set(config.EchoContextKeyCurrentUserID, data["_id"])
			return next(c)
		}
		return util.Response401(c, echo.Map{}, "")
	}
}

// parseToken ...
func parseToken(token string, authSecretKey string) (*jwt.Token, error) {
	if token == "" {
		return nil, fmt.Errorf("Token not found")
	}
	strArr := strings.Split(token, " ")
	if len(strArr) != 2 {
		return nil, fmt.Errorf("Token not found")
	}
	encodedStr := strArr[1]
	tokenData, err := jwt.Parse(encodedStr, func(data *jwt.Token) (interface{}, error) {
		if _, ok := data.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", data.Header["alg"])
		}

		return []byte(authSecretKey), nil
	})
	return tokenData, err
}

// requireLogin to do next action
func requireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userIDString := c.Get(config.EchoContextKeyCurrentUserID).(string)

		if userIDString == "" {
			return util.Response401(c, echo.Map{}, "")
		}

		userID := uuid.FromStringOrNil(userIDString)
		if userID.String() == "" {
			return util.Response401(c, echo.Map{}, "")
		}

		return next(c)
	}
}

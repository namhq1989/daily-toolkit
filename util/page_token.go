package util

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

// PageToken ...
type PageToken struct {
	Page      int64
	Timestamp time.Time
}

func getDefaultPageToken() (response PageToken) {
	response.Page = 0
	response.Timestamp = time.Now()
	return response
}

// PageTokenEncode encode next page token for api response
func PageTokenEncode(page int64, timestamp time.Time) string {
	tokenData := PageToken{
		Page:      page,
		Timestamp: timestamp,
	}
	tokenString, _ := json.Marshal(tokenData)
	encodedString := base64.StdEncoding.EncodeToString(tokenString)
	return encodedString
}

// PageTokenDecode decode page token from query
func PageTokenDecode(c echo.Context) PageToken {
	encodedString := c.QueryParam("pageToken")
	if encodedString == "" {
		return getDefaultPageToken()
	}

	// Decode string
	decoded, err := base64.StdEncoding.DecodeString(encodedString)

	if err != nil {
		return getDefaultPageToken()
	}

	// Parse token
	var pageToken PageToken
	err = json.Unmarshal(decoded, &pageToken)
	if err != nil {
		return getDefaultPageToken()
	}

	return pageToken
}

// PageTokenUsingPage generate next page token using "page"
func PageTokenUsingPage(page int64) string {
	return PageTokenEncode(page, time.Now())
}

// PageTokenUsingTimestamp generate next page token using "timestamp"
func PageTokenUsingTimestamp(timestamp time.Time) string {
	return PageTokenEncode(0, timestamp)
}

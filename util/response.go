package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response ...
type Response struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
}

func sendResponse(c echo.Context, data Response) {
	c.JSON(data.HTTPCode, echo.Map{
		"data":    data.Data,
		"message": data.Message,
	})
}

func getResponse(data interface{}, message string, httpCode int) Response {
	if data == nil {
		data = echo.Map{}
	}

	return Response{
		HTTPCode: httpCode,
		Data:     data,
		Message:  message,
	}
}

// Response200 success
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Thành công!"
	}

	resp := getResponse(data, message, http.StatusOK)
	sendResponse(c, resp)
	return nil
}

// Response400 bad request
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không hợp lệ!"
	}

	resp := getResponse(data, message, http.StatusBadRequest)
	sendResponse(c, resp)
	return nil
}

// Response401 unauthorized
func Response401(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Bạn không có quyền thực hiện hành động này!"
	}

	resp := getResponse(data, message, http.StatusUnauthorized)
	sendResponse(c, resp)
	return nil
}

// Response404 not found
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không tìm thấy!"
	}

	resp := getResponse(data, message, http.StatusNotFound)
	sendResponse(c, resp)
	return nil
}

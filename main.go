package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhq1989/daily-toolkit/route"
)

func main() {
	e := echo.New()

	// Init route
	route.Init(e)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":3000"))
}

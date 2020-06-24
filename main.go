package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/namhq1989/daily-toolkit/route"
)

func main() {
	e := echo.New()

	// Add TLS
	e.Pre(middleware.HTTPSRedirect())
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("daily-toolkit.petprojects.rocks")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	// Init route
	route.Init(e)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.StartAutoTLS(":3000"))
}

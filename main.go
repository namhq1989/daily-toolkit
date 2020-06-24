package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/namhq1989/daily-toolkit/initialize"
	"github.com/namhq1989/daily-toolkit/route"
)

var isRelease = false

func init() {
	initialize.Init()

	// Check env
	env := os.Getenv("ENV")
	isRelease = env == "release"
}

func main() {
	e := echo.New()

	if isRelease {
		// Add TLS
		e.Pre(middleware.HTTPSRedirect())
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("daily-toolkit.petprojects.rocks")
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	}

	// Init route
	route.Init(e)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	// Start server
	if isRelease {
		// Running on port 443, due to TLS required
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(":3000"))
	}
}

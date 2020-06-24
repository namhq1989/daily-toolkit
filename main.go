package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/namhq1989/daily-toolkit/route"
)

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("https://daily-toolkit.petprojects.rocks")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	// Init route
	route.Init(e)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}

package main

import (
	"os"

	"github.com/ilcm96/image-optimize/optimize"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","ip":"${remote_ip}",` +
			`"method":"${method}","uri":"${uri}","status":${status},` +
			`"error":"${error}","latency":"${latency_human}"}` + "\n",
		Output: e.Logger.Output(),
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Static("/", "public")
	e.POST("/optimize", optimize.Optimize)

	e.Logger.Fatal(e.Start(":" + port))
}

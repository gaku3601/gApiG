package main

import (
	"fmt"

	"github.com/gaku3601/gApiG/backendapp/api"
	"github.com/gaku3601/gApiG/backendapp/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Set()
	fmt.Println(config.Value.GetString("env"))

	e := echo.New()
	e.Use(middleware.CORS())

	// Routes
	e.GET("/api", api.Index)
	e.POST("/api", api.Create)

	// Start server
	e.Run(standard.New(":8080"))
}

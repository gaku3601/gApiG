package main

import (
	"github.com/gaku3601/gApiG/backendapp/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/api", api.Index)
	e.POST("/api", api.New)

	// Start server
	e.Run(standard.New(":8093"))
}

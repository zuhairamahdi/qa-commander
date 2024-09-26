package handlers

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/login", Login)
	e.POST("/login", LoginPost)
	// add authenticated routes group
	// g := e.Group("/auth")

}

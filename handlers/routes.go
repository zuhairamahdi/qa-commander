package handlers

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/login", Login)
	e.POST("/login", LoginPost)
	e.GET("/dashboard", Dashboard)
	e.GET("/users", Users)
	// add authenticated routes group
	// g := e.Group("/auth")

}

package handlers

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/login", Login)
}

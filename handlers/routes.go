package handlers

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/login", Login)
	e.POST("/login", LoginPost)
	// TODO paths need to be authenticated
	e.GET("/dashboard", Dashboard)
	e.GET("/users", Users)
	e.GET("/projects", Projects)
	e.GET("/tasks", Tasks)
	e.GET("/reports", Reports)
	// g := e.Group("/auth")

}

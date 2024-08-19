package handlers

import "github.com/labstack/echo/v4"

func Hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

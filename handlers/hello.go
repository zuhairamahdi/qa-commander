package handlers

import (
	"context"
	"qacommander/views"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	component := views.Hello("World")

	// Render the component.
	return component.Render(context.Background(), c.Response().Writer)

}

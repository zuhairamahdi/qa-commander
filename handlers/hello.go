package handlers

import (
	"qacommander/views"

	"github.com/labstack/echo/v4"
)

// Hello renders the Hello component.
func Hello(c echo.Context) error {
	component := views.Hello("World")

	// Render the component.
	return render(c, component)

}

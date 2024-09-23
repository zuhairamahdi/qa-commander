package handlers

import (
	"qacommander/views"

	"github.com/labstack/echo/v4"
)

// Hello renders the Hello component.
func Hello(c echo.Context) error {
	component := views.Hello("World")

	// Create the view properties.
	view := viewProps{
		title:      "Hello",
		includeNav: true,
		activeNav:  "hello",
		c:          c,
		component:  component,
	}
	// Render the component.
	return render(view)

}

package handlers

import (
	"net/http"
	error_pages "qacommander/views/errors"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	// Create the view properties.
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	error_view := error_pages.Error500()
	switch code {
	case http.StatusNotFound:
		error_view = error_pages.Error404()
	case http.StatusInternalServerError:
		error_view = error_pages.Error500()
	case http.StatusForbidden:
		error_view = error_pages.Error403()
	}

	view := viewProps{
		title:      "Error",
		includeNav: false,
		c:          c,
		component:  error_view,
	}
	// Render the component.
	render(view)
}

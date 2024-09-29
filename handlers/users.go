package handlers

import (
	dash "qacommander/views/dashboard"

	"github.com/labstack/echo/v4"
)

func Users(c echo.Context) error {
	component := dash.Dashboard()
	return render(viewProps{
		component:  component,
		title:      "Users",
		includeNav: true,
		c:          c,
	})
}

package handlers

import (
	"fmt"
	dash "qacommander/views/dashboard"

	"github.com/labstack/echo/v4"
)

func Dashboard(c echo.Context) error {
	path := c.Request().URL.Path
	fmt.Println(path)
	component := dash.Dashboard()
	view := viewProps{
		title:      "Dashboard",
		includeNav: true,
		// activeNav:  "login",
		c:         c,
		component: component,
	}
	return render(view)
}

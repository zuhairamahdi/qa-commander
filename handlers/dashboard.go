package handlers

import (
	dash "qacommander/views/dashboard"

	"github.com/labstack/echo/v4"
)

func Dashboard(c echo.Context) error {

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

func Projects(c echo.Context) error {

	component := dash.Projects()
	view := viewProps{
		title:      "Projects",
		includeNav: true,
		// activeNav:  "login",
		c:         c,
		component: component,
	}
	return render(view)
}

func Tasks(c echo.Context) error {

	component := dash.Tasks()
	view := viewProps{
		title:      "Tasks",
		includeNav: true,
		c:          c,
		component:  component,
	}
	return render(view)
}

func Reports(c echo.Context) error {

	component := dash.Reports()
	view := viewProps{
		title:      "Reports",
		includeNav: true,
		c:          c,
		component:  component,
	}
	return render(view)
}

package handlers

import (
	auth_views "qacommander/views/auth"
	dash "qacommander/views/dashboard"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	component := auth_views.Login()
	view := viewProps{
		title:      "Login",
		includeNav: false,
		activeNav:  "login",
		c:          c,
		component:  component,
	}
	return render(view)
}

func LoginPost(c echo.Context) error {
	component := dash.Dashbard()
	view := viewProps{
		title:      "Dashboard",
		includeNav: false,
		// activeNav:  "login",
		c:         c,
		component: component,
	}
	return render(view)
}

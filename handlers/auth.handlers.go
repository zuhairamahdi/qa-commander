package handlers

import (
	"log"
	auth_views "qacommander/views/auth"

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

	// Get the form values.
	username := c.FormValue("username")
	password := c.FormValue("password")
	log.Printf("Username: %s, Password: %s", username, password)

	//redirect to dashboard
	return c.Redirect(302, "/dashboard")
}

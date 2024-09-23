package handlers

import (
	auth_views "qacommander/views/auth"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	component := auth_views.Login()

	return render(c, component)
}

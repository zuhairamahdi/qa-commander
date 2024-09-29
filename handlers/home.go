package handlers

import (
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	return c.Redirect(302, "/dashboard")

}

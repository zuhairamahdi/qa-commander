package handlers

import (
	layout "qacommander/views/layout"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	baseComponent := layout.BaseView(component)
	return baseComponent.Render(c.Request().Context(), c.Response().Writer)
}

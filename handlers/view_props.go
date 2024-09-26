package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type viewProps struct {
	title      string
	includeNav bool
	activeNav  string
	c          echo.Context
	component  templ.Component
}

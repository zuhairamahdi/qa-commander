package handlers

import (
	layout "qacommander/views/layout"
)

func render(view viewProps) error {
	baseComponent := layout.BaseView(view.component)
	return baseComponent.Render(view.c.Request().Context(), view.c.Response().Writer)
}

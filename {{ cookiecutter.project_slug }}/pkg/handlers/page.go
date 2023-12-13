package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/components"
)

// GET /
func HomePage(c echo.Context) error {
	component := components.HomePage(4)
	return component.Render(context.Background(), c.Response().Writer)
}

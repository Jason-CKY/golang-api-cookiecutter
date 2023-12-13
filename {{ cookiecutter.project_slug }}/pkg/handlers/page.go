package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/components"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
)

// GET /
func HomePage(c echo.Context) error {
	_, err := core.GetOrRefreshToken(c)
	if err != nil {
		log.Error(err.Error())
		// return to login page if access token and refresh token both invalid
		component := components.LoginPage()
		return component.Render(context.Background(), c.Response().Writer)
	}
	// show content
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	component := components.HomePage(4)
	return component.Render(context.Background(), c.Response().Writer)
}

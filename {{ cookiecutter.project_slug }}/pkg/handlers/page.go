package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/components"

	{% if cookiecutter.use_oauth %}
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/core"
	{% endif %}
)

// GET /
func HomePage(c echo.Context) error {
	{% if cookiecutter.use_oauth %}
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
	{% else %}
	component := components.HomePage(4)
	return component.Render(context.Background(), c.Response().Writer)
	{% endif %}
}
{% if cookiecutter.use_oauth %}
// GET /login/github
func LoginRedirect(c echo.Context) error {
	// randomly generate this state, then store it in the browser cookies and verify the next token with cookie token
	oauthURL := fmt.Sprintf("%v/login/oauth/authorize?scope=%v&client_id=%v",
		core.GithubHost,
		strings.Join(core.GithubScope, "%20"),
		core.GithubClientID,
	)

	return c.Redirect(http.StatusTemporaryRedirect, oauthURL)
}

// GET /oauth/redirect
func OauthRedirectPage(c echo.Context) error {
	code := c.QueryParam("code")
	code, urlDecodeErr := url.PathUnescape(code)
	if urlDecodeErr != nil {
		log.Errorf("Error decoding code: %v", code)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	oauthResponse, echoHttpErr := core.RequestAccessCode(code)
	if echoHttpErr != nil {
		return echoHttpErr
	}
	// Set access token and refresh token cookie
	c.SetCookie(&http.Cookie{
		Name:  core.AccessTokenCookie,
		Value: oauthResponse.AccessToken,
		Path:  "/", // cookie will be sent to all paths in the same origin
	})
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
{% endif %}

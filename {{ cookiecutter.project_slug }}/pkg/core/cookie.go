package core

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrRefreshToken(c echo.Context) (string, error) {
	var accessToken string
	// check for access token validity
	accessTokenCookie, _ := c.Cookie(AccessTokenCookie)
	// if state cookie is invalid, check for refresh token valdity
	if accessTokenCookie.Valid() != nil {
		// Invalidate cookie
		c.SetCookie(&http.Cookie{
			Name:   AccessTokenCookie,
			Path:   "/", // cookie will be sent to all paths in the same origin
			MaxAge: -1,
		})
		return accessToken, errors.New("access token invalid")

	} else {
		accessToken = accessTokenCookie.Value
	}
	return accessToken, nil
}

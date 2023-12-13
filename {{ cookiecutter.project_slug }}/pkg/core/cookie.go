package core

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetOrRefreshToken(c echo.Context) (string, error) {
	var accessToken string
	// check for access token validity
	accessTokenCookie, _ := c.Cookie(AccessTokenCookie)
	// if state cookie is invalid, check for refresh token valdity
	if accessTokenCookie.Valid() != nil {
		refreshTokenCookie, _ := c.Cookie(RefreshTokenCookie)
		if refreshTokenCookie.Valid() != nil {
			// Invalidate cookie
			c.SetCookie(&http.Cookie{
				Name:   AccessTokenCookie,
				Path:   "/", // cookie will be sent to all paths in the same origin
				MaxAge: -1,
			})
			// Invalidate cookie
			c.SetCookie(&http.Cookie{
				Name:   RefreshTokenCookie,
				Path:   "/", // cookie will be sent to all paths in the same origin
				MaxAge: -1,
			})
			return accessToken, errors.New("access token and Refresh Token both invalid")
		} else {
			// Refresh access token cookie
			// oauthResponse, echoHttpErr := RefreshAccessToken(refreshTokenCookie.Value)
			// if echoHttpErr != nil {
			// 	return accessToken, echoHttpErr
			// }

			// Set access token and refresh token cookie
			c.SetCookie(&http.Cookie{
				Name:   AccessTokenCookie,
				Value:  AccessTokenCookie,
				Path:   "/", // cookie will be sent to all paths in the same origin
				MaxAge: 12,
			})
			c.SetCookie(&http.Cookie{
				Name:  RefreshTokenCookie,
				Value: AccessTokenCookie,
				Path:  "/", // cookie will be sent to all paths in the same origin
				// Set expiry on refresh token even if gitlab doesn't have an expiry for refresh token
				Expires: time.Now().Add(8 * time.Hour), // refresh token set to expire in 8 hours so user has to login again
			})
			accessToken = AccessTokenCookie
		}
	} else {
		accessToken = accessTokenCookie.Value
	}
	return accessToken, nil
}

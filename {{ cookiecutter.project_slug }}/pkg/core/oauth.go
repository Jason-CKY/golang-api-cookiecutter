package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/{{ cookiecutter.author }}/{{ cookiecutter.project_slug }}/pkg/schemas"
)

func RequestAccessCode(authorizationCode string) (schemas.GithubOauthResponse, *echo.HTTPError) {
	endpoint := fmt.Sprintf("%v/login/oauth/access_token", GithubHost)
	reqBody, _ := json.Marshal(map[string]string{
		"client_id":     GithubClientID,
		"client_secret": GithubClientSecret,
		"code":          authorizationCode,
	})
	req, httpErr := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if httpErr != nil {
		log.Error(httpErr.Error())
		return schemas.GithubOauthResponse{}, echo.NewHTTPError(http.StatusBadRequest, httpErr.Error())
	}
	client := &http.Client{}
	res, httpErr := client.Do(req)
	if httpErr != nil {
		log.Error(httpErr.Error())
		return schemas.GithubOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		log.Error(string(body))
		return schemas.GithubOauthResponse{}, echo.NewHTTPError(res.StatusCode, string(body))
	}

	var oauthResponse schemas.GithubOauthResponse
	jsonErr := json.Unmarshal(body, &oauthResponse)
	// error handling for json unmarshaling
	if jsonErr != nil {
		log.Error(string(body))
		log.Error(jsonErr.Error())
		return schemas.GithubOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
	}

	return oauthResponse, nil
}

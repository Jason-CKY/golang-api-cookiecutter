package core

// func RequestAccessCode(authorizationCode string) (schemas.GitlabOauthResponse, *echo.HTTPError) {
// 	endpoint := fmt.Sprintf("%v/oauth/token?client_id=%v&client_secret=%v&code=%v&grant_type=authorization_code&redirect_uri=%v",
// 		GitlabHost,
// 		GitlabClientId,
// 		GitlabClientSecret,
// 		authorizationCode,
// 		RedirectUri)
// 	req, httpErr := http.NewRequest(http.MethodPost, endpoint, nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	if httpErr != nil {
// 		log.Error(httpErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusBadRequest, httpErr.Error())
// 	}
// 	client := &http.Client{}
// 	res, httpErr := client.Do(req)
// 	if httpErr != nil {
// 		log.Error(httpErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
// 	}
// 	defer res.Body.Close()
// 	body, _ := io.ReadAll(res.Body)
// 	if res.StatusCode != 200 {
// 		log.Error(string(body))
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(res.StatusCode, string(body))
// 	}

// 	var oauthResponse schemas.GitlabOauthResponse
// 	jsonErr := json.Unmarshal(body, &oauthResponse)
// 	// error handling for json unmarshaling
// 	if jsonErr != nil {
// 		log.Error(jsonErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
// 	}

// 	return oauthResponse, nil
// }

// func RefreshAccessToken(refreshToken string) (schemas.GitlabOauthResponse, *echo.HTTPError) {
// 	endpoint := fmt.Sprintf("%v/oauth/token?client_id=%v&client_secret=%v&refresh_token=%v&grant_type=refresh_token&redirect_uri=%v",
// 		GitlabHost,
// 		GitlabClientId,
// 		GitlabClientSecret,
// 		refreshToken,
// 		RedirectUri)
// 	req, httpErr := http.NewRequest(http.MethodPost, endpoint, nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	if httpErr != nil {
// 		log.Error(httpErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusBadRequest, httpErr.Error())
// 	}
// 	client := &http.Client{}
// 	res, httpErr := client.Do(req)
// 	if httpErr != nil {
// 		log.Error(httpErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, httpErr.Error())
// 	}
// 	defer res.Body.Close()
// 	body, _ := io.ReadAll(res.Body)
// 	if res.StatusCode != 200 {
// 		log.Error(string(body))
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(res.StatusCode, string(body))
// 	}

// 	var oauthResponse schemas.GitlabOauthResponse
// 	jsonErr := json.Unmarshal(body, &oauthResponse)
// 	// error handling for json unmarshaling
// 	if jsonErr != nil {
// 		log.Error(jsonErr.Error())
// 		return schemas.GitlabOauthResponse{}, echo.NewHTTPError(http.StatusInternalServerError, jsonErr.Error())
// 	}

// 	log.Debugf("Refreshed token. Response: %v", oauthResponse)

// 	return oauthResponse, nil
// }

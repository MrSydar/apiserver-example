package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	auth0Config "mrsydar/apiserver/configs/auth0"
	"mrsydar/apiserver/responses"

	"github.com/labstack/echo/v4"
)

func FetchJWTToken(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return responses.Message(c, http.StatusBadRequest, "code parameter was not provided")
	}

	url := auth0Config.TokenFetchURL
	data := auth0Config.GetDataForTokenFetchWithCode(code)

	response, err := http.PostForm(url, data)
	if err != nil {
		return fmt.Errorf("failed to retrieve JWT token from Auth0 server: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response form Auth0 server: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusForbidden {
			return responses.Message(c, http.StatusForbidden, "got response from auth0: unauthorized")
		} else {
			return fmt.Errorf("got bad response from auth0: %v", err)
		}
	}

	fieldsToCheck := struct {
		Scope string `json:"scope"`
	}{}

	if err := json.Unmarshal(body, &fieldsToCheck); err != nil {
		return fmt.Errorf("failed to unmarshal body for field check: %v", err)
	}

	if !strings.Contains(fieldsToCheck.Scope, "email") {
		return responses.Message(c, http.StatusBadRequest, `"email" scope is required`)
	}

	return c.String(http.StatusOK, string(body))
}

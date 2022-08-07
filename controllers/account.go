package controllers

import (
	"fmt"
	"mrsydar/apiserver/configs/constants/contextnames"
	"mrsydar/apiserver/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccount(c echo.Context) error {
	accountID := c.Get(contextnames.AccountID)

	return responses.Message(c, http.StatusOK, fmt.Sprint(accountID))
}

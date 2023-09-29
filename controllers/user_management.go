package controllers

import (
	"Skripsi/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginUM(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := service.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

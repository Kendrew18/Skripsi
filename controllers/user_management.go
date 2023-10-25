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

func UpdateToken(c echo.Context) error {
	id_user := c.FormValue("id_user")
	Token := c.FormValue("Token")

	result, err := service.Update_Token(id_user, Token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteToken(c echo.Context) error {
	id_user := c.FormValue("id_user")

	result, err := service.Delete_Token(id_user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

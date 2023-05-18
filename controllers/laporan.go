package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InputLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := models.Input_Laporan(id_proyek, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Laporan(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")
	laporan := c.FormValue("total_nilai_kontrak")

	result, err := models.Update_Laporan(id_laporan, laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

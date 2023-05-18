package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InputLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_kontrak := c.FormValue("id_kontrak")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := models.Input_Laporan_Vendor(id_proyek, id_kontrak, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Laporan_Vendor(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")
	laporan := c.FormValue("total_nilai_kontrak")

	result, err := models.Update_Laporan(id_laporan_vendor, laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

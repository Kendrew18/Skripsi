package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InputLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_kontrak := c.FormValue("nomor_kontrak")
	laporan := c.FormValue("total_nilai_kontrak")
	tanggal_laporan := c.FormValue("jenis_pekerjaan")

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

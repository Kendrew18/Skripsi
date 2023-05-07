package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputKontrakVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_user")
	nomor_kontrak := c.FormValue("nomor_kontrak")
	nama_vendor := c.FormValue("nama_vendor")
	total_nilai_kontrak := c.FormValue("total_nilai_kontrak")
	jenis_pekerjaan := c.FormValue("jenis_pekerjaan")
	tanggal_mulai_kontrak := c.FormValue("tanggal_mulai_kontrak")
	tanggal_berakhir_kontrak := c.FormValue("tanggal_berakhir_kontrak")

	tmp, _ := strconv.ParseInt(total_nilai_kontrak, 10, 64)

	result, err := models.Input_Kontrak_Vendor(id_proyek, nomor_kontrak, nama_vendor, tmp,
		jenis_pekerjaan, tanggal_mulai_kontrak, tanggal_berakhir_kontrak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadKontrakVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_user")

	result, err := models.Read_Kontrak_Vendor(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

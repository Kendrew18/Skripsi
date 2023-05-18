package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputKontrakVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	nomor_kontrak := c.FormValue("nomor_kontrak")
	nama_vendor := c.FormValue("nama_vendor")
	total_nilai_kontrak := c.FormValue("total_nilai_kontrak")
	jenis_pekerjaan_vendor := c.FormValue("jenis_pekerjaan_vendor")
	pekerjaan_vendor := c.FormValue("pekerjaan_vendor")
	tanggal_mulai_kontrak := c.FormValue("tanggal_mulai_kontrak")
	tanggal_berakhir_kontrak := c.FormValue("tanggal_berakhir_kontrak")
	tanggal_pengiriman := c.FormValue("tanggal_pengiriman")
	tanggal_pengerjaan_dimulai := c.FormValue("tanggal_pengerjaan_dimulai")
	tanggal_pengerjaan_selesai := c.FormValue("tanggal_pengerjaan_selesai")

	tmp, _ := strconv.ParseInt(total_nilai_kontrak, 10, 64)

	result, err := models.Input_Kontrak_Vendor(id_proyek, nomor_kontrak, nama_vendor, tmp,
		jenis_pekerjaan_vendor, pekerjaan_vendor, tanggal_mulai_kontrak, tanggal_berakhir_kontrak,
		tanggal_pengiriman, tanggal_pengerjaan_dimulai, tanggal_pengerjaan_selesai)

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

func DeleteKontrakVendor(c echo.Context) error {
	id_kontrak := c.FormValue("id_kontrak")

	result, err := models.Delete_Kontrak_Vendor(id_kontrak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

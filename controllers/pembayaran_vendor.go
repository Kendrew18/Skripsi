package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputPembayaranVendor(c echo.Context) error {
	id_kontrak := c.FormValue("id_kontrak")
	nomor_invoice := c.FormValue("nomor_invoice")
	jumlah_pembayaran := c.FormValue("jumlah_pembayaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")

	tmp, _ := strconv.ParseInt(jumlah_pembayaran, 10, 64)

	result, err := models.Input_Pembayaran_Vendor(id_kontrak, nomor_invoice, tmp, tanggal_pembayaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadPembayaranVendor(c echo.Context) error {
	id_kontrak := c.FormValue("id_kontrak")

	result, err := models.Read_Pembayaran_Vendor(id_kontrak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Foto_Invoice(c echo.Context) error {
	path := c.FormValue("path")
	return c.File(path)
}

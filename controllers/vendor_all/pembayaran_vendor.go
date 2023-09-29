package vendor_all

import (
	"Skripsi/service/vendor_all"
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

	result, err := vendor_all.Input_Pembayaran_Vendor(id_kontrak, nomor_invoice, tmp, tanggal_pembayaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadPembayaranVendor(c echo.Context) error {
	id_kontrak := c.FormValue("id_kontrak")

	result, err := vendor_all.Read_Pembayaran_Vendor(id_kontrak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UploadInvoice(c echo.Context) error {
	id_PV := c.FormValue("id_PV")

	result, err := vendor_all.Upload_Invoice(id_PV, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadFotoPembayaranvendor(c echo.Context) error {
	id_PV := c.FormValue("id_PV")

	result, err := vendor_all.Read_Foto_Pembayaran_vendor(id_PV)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadFoto(c echo.Context) error {
	path := c.FormValue("path")
	return c.File(path)
}

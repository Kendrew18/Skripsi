package vendor_all

import (
	"Skripsi/models/vendor_all"
	"github.com/labstack/echo/v4"
	"net/http"
)

//data
func InputLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_kontrak := c.FormValue("id_kontrak")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := vendor_all.Input_Laporan_Vendor(id_proyek, id_kontrak, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := vendor_all.Read_Laporan_Vendor(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")
	id_kontrak := c.FormValue("id_kontrak")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := vendor_all.Update_Laporan_Vendor(id_laporan_vendor, id_kontrak, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStatusLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Update_Status_Laporan_Vendor(id_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//foto
func UploadFotolaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Upload_Foto_laporan_vendor(id_laporan_vendor, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadFotolaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Read_Foto_Laporan_Vendor(id_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
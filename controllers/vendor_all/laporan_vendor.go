package vendor_all

import (
	"Skripsi/service/vendor_all"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Input_Laporan_Vendor
func InputLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_kontrak := c.FormValue("id_kontrak")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")
	check_Box := c.FormValue("check_Box")

	result, err := vendor_all.Input_Laporan_Vendor(id_proyek, laporan, tanggal_laporan, id_kontrak, check_Box)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Laporan_Vendor
func ReadLaporanVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := vendor_all.Read_Laporan_Vendor(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Laporan_Vendor
func UpdateLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")
	id_kontrak := c.FormValue("id_kontrak")
	laporan := c.FormValue("laporan")
	check_Box := c.FormValue("check_Box")

	result, err := vendor_all.Update_Laporan_Vendor(id_laporan_vendor, laporan, id_kontrak, check_Box)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Status_Laporan_Vendor
func UpdateStatusLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Update_Status_Laporan_Vendor(id_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Upload_Foto_Laporan_Vendor
func UploadFotolaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Upload_Foto_laporan_vendor(id_laporan_vendor, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Foto_Laporan_Vendor
func ReadFotolaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Read_Foto_Laporan_Vendor(id_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_Task_Vendor
func SeeTaskVendor(c echo.Context) error {
	tanggal_laporan_vendor := c.FormValue("tanggal_laporan_vendor")

	result, err := vendor_all.See_Task_Vendor(tanggal_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_laporan_Vendor
func DeleteLaporanVendor(c echo.Context) error {
	id_laporan_vendor := c.FormValue("id_laporan_vendor")

	result, err := vendor_all.Delete_laporan_Vendor(id_laporan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

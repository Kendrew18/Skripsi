package vendor_all

import (
	"Skripsi/service/vendor_all"
	"github.com/labstack/echo/v4"
	"net/http"
)

//input vendor
func InputVendor(c echo.Context) error {
	nama_vendor := c.FormValue("nama_vendor")
	pekerjaan_vendor := c.FormValue("pekerjaan_vendor")

	result, err := vendor_all.Input_Vendor(nama_vendor, pekerjaan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read vendor
func ReadVendor(c echo.Context) error {

	result, err := vendor_all.Read_Vendor()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//edit vendor
func EditVendor(c echo.Context) error {
	id_vendor := c.FormValue("id_vendor")
	nama_vendor := c.FormValue("nama_vendor")
	pekerjaan_vendor := c.FormValue("pekerjaan_vendor")

	result, err := vendor_all.Edit_Vendor(id_vendor, nama_vendor, pekerjaan_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//delete vendor
func DeleteVendor(c echo.Context) error {
	id_vendor := c.FormValue("id_vendor")

	result, err := vendor_all.Delete_Vendor(id_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

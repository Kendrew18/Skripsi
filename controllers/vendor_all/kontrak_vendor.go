package vendor_all

import (
	"Skripsi/service/vendor_all"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Kontrak_Vendor
func InputKontrakVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_master_vendor := c.FormValue("id_master_vendor")
	total_nilai_kontrak := c.FormValue("total_nilai_kontrak")
	tanggal_mulai_kontrak := c.FormValue("tanggal_mulai_kontrak")
	tanggal_berakhir_kontrak := c.FormValue("tanggal_berakhir_kontrak")
	tanggal_pengiriman := c.FormValue("tanggal_pengiriman")
	tanggal_pengerjaan_dimulai := c.FormValue("tanggal_pengerjaan_dimulai")
	tanggal_pengerjaan_selesai := c.FormValue("tanggal_pengerjaan_selesai")

	tmp, _ := strconv.ParseInt(total_nilai_kontrak, 10, 64)

	result, err := vendor_all.Input_Kontrak_Vendor(id_proyek, id_master_vendor, tmp,
		tanggal_mulai_kontrak, tanggal_berakhir_kontrak, tanggal_pengiriman,
		tanggal_pengerjaan_dimulai, tanggal_pengerjaan_selesai)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Kontrak_Vendor
func ReadKontrakVendor(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_master_vendor := c.FormValue("id_master_vendor")

	result, err := vendor_all.Read_Kontrak_Vendor(id_proyek, id_master_vendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_Kontrak_Vendor
func DeleteKontrakVendor(c echo.Context) error {
	id_kontrak := c.FormValue("id_kontrak")

	result, err := vendor_all.Delete_Kontrak_Vendor(id_kontrak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Pick_Vendor
func PickVendor(c echo.Context) error {

	result, err := vendor_all.Pick_Vendor()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Data_Filter
func DataFilter(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	result, err := vendor_all.Data_Filter(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

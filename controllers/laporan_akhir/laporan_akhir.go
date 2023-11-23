package laporan_akhir

import (
	"Skripsi/service/laporan"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//update_status_penawaran
func ReadLaporanAkhir(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := laporan.Read_Laporan_Akhir(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update_judul_penawaran
func UpdateStatus(c echo.Context) error {
	id_laporan_akhir := c.FormValue("id_laporan_akhir")
	status, _ := strconv.Atoi(c.FormValue("status"))

	result, err := laporan.Update_Status(status, id_laporan_akhir)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

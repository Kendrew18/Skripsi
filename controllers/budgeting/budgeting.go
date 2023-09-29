package budgeting

import (
	"Skripsi/service/budgeting"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Realisasi
func InputRealisasi(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

	nm, _ := strconv.ParseInt(nominal_pembayaran, 10, 64)

	result, err := budgeting.Input_Realisasi(id_proyek, id_sub_pekerjaan, id_kontrak,
		perihal_pengeluaran, tanggal_pembayaran, nm, catatan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Realisasi
func ReadRealisasi(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")

	result, err := budgeting.Read_Realisasi(id_proyek, id_sub_pekerjaan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_Realisasi
func DeleteRealisasi(c echo.Context) error {
	id_realisasi := c.FormValue("id_realisasi")

	result, err := budgeting.Delete_Realisasi(id_realisasi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Realisasi
func UpdateRealisasi(c echo.Context) error {
	id_realisasi := c.FormValue("id_realisasi")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

	nm, _ := strconv.ParseInt(nominal_pembayaran, 10, 64)

	result, err := budgeting.Update_Realisasi(id_realisasi, id_kontrak, perihal_pengeluaran,
		tanggal_pembayaran, nm, catatan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Budgeting
func ReadBudgeting(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := budgeting.Read_Budgeting(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

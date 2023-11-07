package budgeting

import (
	"Skripsi/service/budgeting"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Realisasi
func InputDetailBudgeting(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

	nm, _ := strconv.ParseInt(nominal_pembayaran, 10, 64)

	result, err := budgeting.Input_Detail_Budgeting(id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, nm, catatan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Detail_Budgeting
func ReadDetailBudgeting(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	id_laporan := c.FormValue("id_laporan")

	result, err := budgeting.Read_Detail_Budgeting(id_proyek, id_sub_pekerjaan, id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_Detail_Budgeting
func DeleteDetailBudgeting(c echo.Context) error {
	id_budgeting := c.FormValue("id_budgeting")

	result, err := budgeting.Delete_Detail_Budgeting(id_budgeting)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update_Detail_Budgeting
func UpdateDetailBudgeting(c echo.Context) error {
	id_budgeting := c.FormValue("id_budgeting")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

	nm, _ := strconv.ParseInt(nominal_pembayaran, 10, 64)

	result, err := budgeting.Update_Detail_Budgeting(id_budgeting, id_kontrak, perihal_pengeluaran, nm, catatan)

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

//Pilih_Kontrak
func PilihKontrak(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := budgeting.Pilih_Kontrak(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

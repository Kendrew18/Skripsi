package analisa_budgeting

import (
	"Skripsi/service/analisa_budgeting"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadAnalisaBudgeting(c echo.Context) error {
	tanggal_sekarang := c.FormValue("tanggal_sekarang")
	id_proyek := c.FormValue("id_proyek")

	result, err := analisa_budgeting.Analisa_Budgeting(tanggal_sekarang, id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

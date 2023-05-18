package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

//create
func InputProyek(c echo.Context) error {
	id_user := c.FormValue("id_user")
	nama_proyek := c.FormValue("nama_proyek")
	jumlah_lantai := c.FormValue("jumlah_lantai")
	luas_tanah := c.FormValue("luas_tanah")
	nama_penanggungjawab_proyek := c.FormValue("nama")

	result, err := models.Input_Proyek(id_user, nama_proyek, jumlah_lantai, luas_tanah, nama_penanggungjawab_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Nama_Proyek
func ReadNamaProyek(c echo.Context) error {
	result, err := models.Read_Nama_Proyek()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadNamaProyekHistory(c echo.Context) error {
	result, err := models.Read_Nama_Proyek_history()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Proyek
func ReadProyek(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	result, err := models.Read_Proyek(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadHistory(c echo.Context) error {
	result, err := models.Read_History()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update Status
func FinishProyek(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Finish_Proyek(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

package jadwal

import (
	"Skripsi/models/jadwal"
	"github.com/labstack/echo/v4"
	"net/http"
)

//data
func InputLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := jadwal.Input_Laporan(id_proyek, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := jadwal.Read_Laporan(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := jadwal.Update_Laporan(id_laporan, laporan, tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStatusLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Update_Status_Laporan(id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//foto
func UploadFotolaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Upload_Foto_laporan(id_laporan, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadFotolaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Read_Foto_Laporan(id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

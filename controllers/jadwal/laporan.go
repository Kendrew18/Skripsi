package jadwal

import (
	"Skripsi/service/jadwal"
	"github.com/labstack/echo/v4"
	"net/http"
)

//input-laporan
func InputLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")
	id_penjadwalan := c.FormValue("id_penjadwalan")
	check := c.FormValue("check")

	result, err := jadwal.Input_Laporan(id_proyek, laporan, tanggal_laporan,
		id_penjadwalan, check)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read-laporan
func ReadLaporan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := jadwal.Read_Laporan(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update-laporan
func UpdateLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")
	laporan := c.FormValue("laporan")
	id_penjadwalan := c.FormValue("id_penjadwalan")
	check := c.FormValue("check")

	result, err := jadwal.Update_Laporan(id_laporan, laporan, id_penjadwalan, check)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update-status-laporan
func UpdateStatusLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Update_Status_Laporan(id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_Laporan
func DeleteLaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Delete_Laporan(id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//upload-foto-laporan
func UploadFotolaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Upload_Foto_Laporan(id_laporan, c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read-foto-laporan
func ReadFotolaporan(c echo.Context) error {
	id_laporan := c.FormValue("id_laporan")

	result, err := jadwal.Read_Foto_Laporan(id_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_Task
func SeeTask(c echo.Context) error {
	tanggal_laporan := c.FormValue("tanggal_laporan")

	result, err := jadwal.See_Task(tanggal_laporan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

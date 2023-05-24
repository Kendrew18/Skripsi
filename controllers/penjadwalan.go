package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//input tanggal mulai
func InputTanggalMulai(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	tanggal := c.FormValue("tanggal")
	result, err := models.Input_Tanggal_Mulai(id_proyek, tanggal)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read tanggal mulai
func ReadTanggalMulai(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Tanggal_Mulai(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read judul penawaran
func ReadJudulPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Judul_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//input task penjadwalan
func InputTaskPenjadwalan(c echo.Context) error {
	id_penawaran := c.FormValue("id_penawaran")
	id_proyek := c.FormValue("id_proyek")
	nama_task := c.FormValue("nama_task")
	waktu_optimis := c.FormValue("waktu_optimis")
	waktu_pesimis := c.FormValue("waktu_pesimis")
	waktu_realistis := c.FormValue("waktu_realistis")

	WO, _ := strconv.ParseFloat(waktu_optimis, 64)
	WP, _ := strconv.ParseFloat(waktu_pesimis, 64)
	WR, _ := strconv.ParseFloat(waktu_realistis, 64)

	result, err := models.Input_Task_Penjadwalan(id_penawaran, id_proyek, nama_task, WO, WP, WR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read task
func ReadTask(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")

	result, err := models.Read_Task(id_proyek, id_penawaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//input dependentcies
func Inputdepedentcies(c echo.Context) error {
	id_jadwal := c.FormValue("id_jadwal")
	depedentcies := c.FormValue("depedentcies")

	result, err := models.Input_Dependentcies(id_jadwal, depedentcies)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//generate jadwal
func GenerateJadwal(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	result, err := models.Generate_Jadwal(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read_jadwal
func ReadJadwal(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")

	result, err := models.Read_Jadwal(id_proyek, id_penawaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadDep(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_dep(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

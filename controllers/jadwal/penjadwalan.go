package jadwal

import (
	"Skripsi/models/jadwal"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//input durasi (done)
func InputDurasitask(c echo.Context) error {
	id_penjadwalan := c.FormValue("id_penjadwalan")
	waktu_optimis := c.FormValue("waktu_optimis")
	waktu_pesimis := c.FormValue("waktu_pesimis")
	waktu_realistis := c.FormValue("waktu_realistis")

	WO, _ := strconv.ParseFloat(waktu_optimis, 64)
	WP, _ := strconv.ParseFloat(waktu_pesimis, 64)
	WR, _ := strconv.ParseFloat(waktu_realistis, 64)

	result, err := jadwal.Input_Durasi_task(id_penjadwalan, WO, WP, WR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read task
func ReadTask(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := jadwal.Read_Task(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Depedentcies
func ReadDep(c echo.Context) error {
	id_penjadwalan := c.FormValue("id_penjadwalan")
	id_proyek := c.FormValue("id_proyek")

	result, err := jadwal.Read_dep(id_proyek, id_penjadwalan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//input dependentcies
func Inputdepedentcies(c echo.Context) error {
	id_jadwal := c.FormValue("id_jadwal")
	depedentcies := c.FormValue("depedentcies")

	result, err := jadwal.Input_Dependentcies(id_jadwal, depedentcies)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//generate jadwal
func GenerateJadwal(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	result, err := jadwal.Generate_Jadwal(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read_jadwal
func ReadJadwal(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := jadwal.Read_Jadwal(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Edit_Durasi_Tanggal
func EditDurTgl(c echo.Context) error {
	id_penjadwalan := c.FormValue("id_penjadwalan")
	tanggal_pekerjaan_mulai := c.FormValue("tanggal_pekerjaan_mulai")
	durasi := c.FormValue("durasi")

	drs, _ := strconv.Atoi(durasi)

	result, err := jadwal.Edit_Dur_Tgl(id_penjadwalan, tanggal_pekerjaan_mulai, drs)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_Calender_All
func SeeCalenderAll(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	status_user := c.FormValue("status_user")

	su, _ := strconv.Atoi(status_user)

	result, err := jadwal.See_Calender(id_proyek, su)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

package proyek

import (
	"Skripsi/models/proyek"
	"github.com/labstack/echo/v4"
	"net/http"
)

//Input_Proyek
func InputProyek(c echo.Context) error {
	id_user := c.FormValue("id_user")
	nama_proyek := c.FormValue("nama_proyek")
	jumlah_lantai := c.FormValue("jumlah_lantai")
	luas_tanah := c.FormValue("luas_tanah")
	nama_penanggungjawab_proyek := c.FormValue("nama_penanggungjawab")
	nama_client := c.FormValue("nama_client")
	jenis_gedung := c.FormValue("jenis_gedung")
	alamat := c.FormValue("alamat")
	tanggal_mulai_kerja := c.FormValue("tanggal_mulai_kerja")

	result, err := proyek.Input_Proyek(id_user, nama_proyek, nama_client, jenis_gedung,
		alamat, luas_tanah, jumlah_lantai, nama_penanggungjawab_proyek, tanggal_mulai_kerja)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Nama_Proyek
func ReadNamaProyek(c echo.Context) error {
	result, err := proyek.Read_Nama_Proyek()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadNamaProyekHistory(c echo.Context) error {
	result, err := proyek.Read_Nama_Proyek_history()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Proyek
func ReadProyek(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	result, err := proyek.Read_Proyek(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadHistory(c echo.Context) error {
	result, err := proyek.Read_History()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Update Status
func FinishProyek(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := proyek.Finish_Proyek(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

package tagihan

import (
	"Skripsi/service/tagihan"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//Input_Tagihan
func InputTagihan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	perihal := c.FormValue("perihal")
	tanggal_pemberian_kwitansi := c.FormValue("tanggal_pemberian_kwitansi")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_keseluruhan := c.FormValue("nominal_keseluruhan")
	id_penawaran := c.FormValue("id_penawaran")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	nominal := c.FormValue("nominal")

	nk, _ := strconv.ParseInt(nominal_keseluruhan, 10, 64)

	result, err := tagihan.Input_Tagihan(id_proyek, perihal, tanggal_pemberian_kwitansi,
		tanggal_pembayaran, nk, id_penawaran, id_sub_pekerjaan, nominal)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Read_Realisasi
func ReadRealisasi(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := tagihan.Read_Realisasi(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Delete_Tagihan
func DeleteTagihan(c echo.Context) error {
	id_tagihan := c.FormValue("id_tagihan")

	result, err := tagihan.Delete_Tagihan(id_tagihan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_Judul
func SeeJudul(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := tagihan.See_Judul(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//See_Sub_Pekerjaan
func SeeSubPekerjaan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")

	result, err := tagihan.See_Sub_Pekerjaan(id_proyek, id_penawaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

package penawaran

import (
	"Skripsi/models/penawaran"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

//input_penawaran
func InputPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	judul := c.FormValue("judul")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")

	result, err := penawaran.Input_Penawaran(id_proyek, judul, sub_pekerjaan, catatan, jumlah, satuan, harga, sub_total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Input_Sub_Pekerjaan
func InputSubPekerjaan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")

	jmlh, _ := strconv.ParseFloat(jumlah, 64)
	hrg, _ := strconv.ParseInt(harga, 10, 64)

	sbt_f, _ := strconv.ParseFloat(sub_total, 64)

	result, err := penawaran.Input_Sub_Pekerjaan(id_proyek, id_penawaran, sub_pekerjaan, catatan, jmlh, satuan, hrg, sbt_f)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//read_penawaran
func ReadPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := penawaran.Read_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update_status_penawaran
func UpdateStatusPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := penawaran.Update_Status_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update_judul_penawaran
func UpdateJudulPenawaran(c echo.Context) error {
	id_penawaran := c.FormValue("id_penawaran")
	judul := c.FormValue("judul")

	result, err := penawaran.Update_Judul_Penawaran(id_penawaran, judul)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//update_item_penawaran
func UpdateItemPenawaran(c echo.Context) error {
	id_penawaran := c.FormValue("id_penawaran")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")

	jm, _ := strconv.ParseFloat(jumlah, 64)
	hg, _ := strconv.ParseInt(harga, 10, 64)
	tt, _ := strconv.ParseFloat(sub_total, 64)

	result, err := penawaran.Update_Item_Penawaran(id_penawaran, id_sub_pekerjaan, sub_pekerjaan, catatan, jm, satuan, hg, tt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//InputTambahanSubPekerjaan
func InputTambahanSubPekerjaan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")
	tanggal_pekerjaan_dimulai := c.FormValue("tanggal_pekerjaan_dimulai")
	durasi := c.FormValue("durasi")

	dr, _ := strconv.Atoi(durasi)

	sbt_f, _ := strconv.ParseFloat(sub_total, 64)

	jmlh, _ := strconv.ParseFloat(jumlah, 64)
	hrg, _ := strconv.ParseInt(harga, 10, 64)

	result, err := penawaran.Input_Tambahan_Sub_Pekerjaan(id_proyek, id_penawaran, sub_pekerjaan, catatan, jmlh,
		satuan, hrg, sbt_f, tanggal_pekerjaan_dimulai, dr)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//InputTambahanPekerjaanTambah
func InputTambahanPekerjaanTambah(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	judul := c.FormValue("judul")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")
	tanggal_mulai := c.FormValue("tanggal_mulai")
	durasi := c.FormValue("durasi")

	result, err := penawaran.Input_Tambahan_Pekerjaan_Tambah(id_proyek, judul, sub_pekerjaan, catatan, jumlah,
		satuan, harga, sub_total, tanggal_mulai, durasi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//Pilih_Judul_Pekerjaan
func PilihJudulPekerjaan(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := penawaran.Pilih_Judul_Pekerjaan(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

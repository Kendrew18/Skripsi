package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InputHeaderPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	kode_surat := c.FormValue("kode_surat")
	tanggal_dibuat := c.FormValue("id_header_penawaran")
	nama_perusahaan := c.FormValue("nama_perusahaan")
	alamat_perusahaan := c.FormValue("alamat_perusahaan")

	result, err := models.Input_Header_Penawaran(id_proyek, kode_surat, tanggal_dibuat, nama_perusahaan, alamat_perusahaan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadHeaderPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Header_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InputPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	judul := c.FormValue("judul")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	keterangan := c.FormValue("keterangan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	total := c.FormValue("total")

	result, err := models.Input_Penawaran(id_proyek, judul, sub_pekerjaan, keterangan, jumlah, satuan, harga, total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ReadPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Read_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStatusPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")

	result, err := models.Update_Status_Penawaran(id_proyek)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateHeaderPenawaran(c echo.Context) error {
	id_proyek := c.FormValue("id_proyek")
	kode_surat := c.FormValue("kode_surat")
	tanggal_dibuat := c.FormValue("id_header_penawaran")
	nama_perusahaan := c.FormValue("nama_perusahaan")
	alamat_perusahaan := c.FormValue("alamat_perusahaan")

	result, err := models.Update_Header_Penawaran(id_proyek, kode_surat, tanggal_dibuat, nama_perusahaan, alamat_perusahaan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateJudulPenawaran(c echo.Context) error {
	id_penawaran := c.FormValue("id_penawaran")
	judul := c.FormValue("judul")

	result, err := models.Update_Judul_Penawaran(id_penawaran, judul)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateItemPenawaran(c echo.Context) error {
	id_penawaran := c.FormValue("id_penawaran")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	keterangan := c.FormValue("keterangan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	total := c.FormValue("total")

	jm, _ := strconv.ParseFloat(jumlah, 64)
	hg, _ := strconv.ParseInt(harga, 10, 64)
	tt, _ := strconv.ParseInt(total, 10, 64)

	result, err := models.Update_Item_Penawaran(id_penawaran, id_sub_pekerjaan, sub_pekerjaan, keterangan, jm, satuan, hg, tt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

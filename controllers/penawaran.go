package controllers

import (
	"Skripsi/models"
	"github.com/labstack/echo/v4"
	"net/http"
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

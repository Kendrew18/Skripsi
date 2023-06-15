package routes

import (
	controllers "Skripsi/controllers"
	"Skripsi/controllers/jadwal"
	"Skripsi/controllers/penawaran"
	"Skripsi/controllers/proyek"
	"Skripsi/controllers/vendor_all"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	um := e.Group("/um")
	pryk := e.Group("/pryk")
	kv := e.Group("/kv")
	pv := e.Group("/pv")
	pen := e.Group("/pen")
	LV := e.Group("/LV")
	LP := e.Group("/LP")
	PJDL := e.Group("/PJDL")
	FT := e.Group("/FT")

	//MAIN
	//User_Management
	um.GET("/login-user", controllers.LoginUM)

	//Proyek(done)
	//Create
	pryk.POST("/input-proyek", proyek.InputProyek)
	//Read Nama Proyek
	pryk.GET("/Read-Nama", proyek.ReadNamaProyek)
	//Read Detail
	pryk.GET("/Read-proyek", proyek.ReadProyek)
	//Finish (bag 2)
	pryk.PUT("/finish-proyek", proyek.FinishProyek)
	//History Proyek (bag 2)
	//Read Nama Proyek
	pryk.GET("/Read-Nama-his", proyek.ReadNamaProyekHistory)
	//Read Detail
	pryk.GET("/Read-proyek-his", proyek.ReadHistory)

	//Edit (bag 2)

	//Kontrak-Vendor
	//input
	kv.POST("/input-kv", vendor_all.InputKontrakVendor)
	//read
	kv.GET("/read-kv", vendor_all.ReadKontrakVendor)
	//delete
	kv.DELETE("/delete-kontrak", vendor_all.DeleteKontrakVendor)

	//Pembayaran-Vendor
	//input
	pv.POST("/input-pv", vendor_all.InputPembayaranVendor)
	//read
	pv.GET("/read-pv", vendor_all.ReadPembayaranVendor)
	//upload-foto-invoice
	pv.POST("/upload-fi", vendor_all.UploadInvoice)
	//Read-foto-laporan
	pv.GET("/read-path-foto", vendor_all.ReadFotoPembayaranvendor)

	//laporan-vendor
	//Create
	LV.POST("/input-lv", vendor_all.InputLaporanVendor)
	//read
	LV.GET("/Read-lv", vendor_all.ReadLaporanVendor)
	//update
	LV.PUT("/update-lv", vendor_all.UpdateLaporanVendor)
	//upload foto
	LV.POST("/upload-foto", vendor_all.UploadFotolaporanVendor)
	//Read-foto-laporan
	LV.GET("/read-path-foto", vendor_all.ReadFotolaporanVendor)
	//update_status_laporan_vendor
	LV.PUT("/update-status", vendor_all.UpdateStatusLaporanVendor)
	//See Kontrak Vendor
	LV.GET("/see-kontak-vendor", vendor_all.SeeTaskVendor)

	//Penawaran (done)
	//create
	pen.POST("/input-pen", penawaran.InputPenawaran)
	//input_sub_penawaran
	pen.POST("/input-sub-pen", penawaran.InputSubPekerjaan)
	//read
	pen.GET("/read-pen", penawaran.ReadPenawaran)
	//Update Status Penawaran
	pen.PUT("/update-status", penawaran.UpdateStatusPenawaran)
	//update_judul
	pen.PUT("/update-judul", penawaran.UpdateJudulPenawaran)
	//update item
	pen.PUT("/update-item", penawaran.UpdateItemPenawaran)
	//Input_Tambahan_Sub_Pekerjaan
	pen.POST("/input-tambahan-sub-pekerjaan", penawaran.InputTambahanSubPekerjaan)
	//Input_Tambahan_Pekerjaan_Tambah
	pen.POST("input-tambahan-pekerjaan-tambah", penawaran.InputTambahanPekerjaanTambah)
	//pilih_judul_pekerjaan
	pen.GET("pilih-judul-pekerjaan", penawaran.PilihJudulPekerjaan)

	//Penjadwalan
	//Input-Durasi-Task
	PJDL.POST("/input-durasi-task", jadwal.InputDurasitask)
	//read task
	PJDL.GET("/read-task", jadwal.ReadTask)
	//read_dep
	PJDL.GET("/read-dep", jadwal.ReadDep)
	//add dependentcies
	PJDL.PUT("/input-depedentcies", jadwal.Inputdepedentcies)
	//generate jadwal cpm
	PJDL.PUT("/generate-jadwal", jadwal.GenerateJadwal)
	//Read_Jadwal
	PJDL.GET("/read-jadwal", jadwal.ReadJadwal)
	//Edit-Durasi-Tanggal
	PJDL.PUT("/edit-rur-tgl", jadwal.EditDurTgl)
	//See-Calender-All
	PJDL.GET("see-calender-all", jadwal.SeeCalenderAll)

	//Laporan
	//create
	LP.POST("/input-lp", jadwal.InputLaporan)
	//read
	LP.GET("/read-lp", jadwal.ReadLaporan)
	//update
	LP.PUT("/update-lp", jadwal.UpdateLaporan)
	//upload-foto-laporan
	LP.POST("upload-foto-laporan", jadwal.UploadFotolaporan)
	//Read-foto-laporan
	LP.GET("/read-path-foto", jadwal.ReadFotolaporan)
	//update_status_laporan
	LP.PUT("/update-status", jadwal.UpdateStatusLaporan)
	//See Task
	LP.GET("/see-task", jadwal.SeeTask)

	//foto
	//get image foto
	FT.GET("/read-foto", vendor_all.ReadFoto)

	//PENDUKUNG//

	//Budgeting
	//input
	//update
	//read

	//kwitansi-pembayaran

	//Analisa Budgeting(perkiraan budget mingguan)
	//Read

	return e
}

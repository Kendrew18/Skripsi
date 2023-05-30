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
	ph := e.Group("/ph")
	pen := e.Group("/pen")
	LV := e.Group("/LV")
	LP := e.Group("/LP")
	PJDL := e.Group("/PJDL")
	FT := e.Group("/FT")

	//MAIN
	//User_Management
	um.GET("/login-user", controllers.LoginUM)

	//Proyek
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

	//Penawaran(header)
	//create
	ph.POST("/input-ph", penawaran.InputHeaderPenawaran)
	//read
	ph.GET("/read-ph", penawaran.ReadHeaderPenawaran)
	//update

	//Penawaran
	//create
	pen.POST("/input-pen", penawaran.InputPenawaran)
	//read
	pen.GET("/read-pen", penawaran.ReadPenawaran)
	//Update Status Penawaran
	pen.PUT("/update-status", penawaran.UpdateStatusPenawaran)
	//update kop
	pen.PUT("/UpdateHeaderPenawaran", penawaran.UpdateHeaderPenawaran)
	//update_judul
	pen.PUT("/update-judul", penawaran.UpdateJudulPenawaran)
	//update item
	pen.PUT("/update-item", penawaran.UpdateItemPenawaran)

	//Penjadwalan
	//input tanggal mulai
	PJDL.PUT("/tgl-ml", jadwal.InputTanggalMulai)
	//read tanggal mulai
	PJDL.GET("/read-tgl-ml", jadwal.ReadTanggalMulai)
	//read_judul_penawaran
	PJDL.GET("/read-judul-penawaran", jadwal.ReadJudulPenawaran)
	//create task
	PJDL.POST("/input-task-penjadwalan", jadwal.InputTaskPenjadwalan)
	//read task
	PJDL.GET("/read-task", jadwal.ReadTask)
	//add dependentcies
	PJDL.PUT("/input-depedentcies", jadwal.Inputdepedentcies)
	//generate jadwal cpm
	PJDL.PUT("/generate-jadwal", jadwal.GenerateJadwal)
	//Read_Jadwal
	PJDL.GET("/read-jadwal", jadwal.ReadJadwal)
	//read_dep
	PJDL.GET("/read-dep", jadwal.ReadDep)
	//edit

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

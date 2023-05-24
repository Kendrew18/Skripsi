package routes

import (
	controllers "Skripsi/controllers"
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
	pryk.POST("/input-proyek", controllers.InputProyek)

	//Read Nama Proyek
	pryk.GET("/Read-Nama", controllers.ReadNamaProyek)

	//Read Detail
	pryk.GET("/Read-proyek", controllers.ReadProyek)

	//Finish (bag 2)
	pryk.PUT("/finish-proyek", controllers.FinishProyek)

	//History Proyek (bag 2)
	//Read Nama Proyek
	pryk.GET("/Read-Nama-his", controllers.ReadNamaProyekHistory)

	//Read Detail
	pryk.GET("/Read-proyek-his", controllers.ReadHistory)

	//Edit (bag 2)

	//Kontrak-Vendor
	//input
	kv.POST("/input-kv", controllers.InputKontrakVendor)
	//read
	kv.GET("/read-kv", controllers.ReadKontrakVendor)
	//delete
	kv.DELETE("/delete-kontrak", controllers.DeleteKontrakVendor)

	//Pembayaran-Vendor
	//input
	pv.POST("/input-pv", controllers.InputPembayaranVendor)
	//read
	pv.GET("/read-pv", controllers.ReadPembayaranVendor)
	//upload-foto-invoice
	pv.POST("/upload-fi", controllers.UploadInvoice)
	//Read-foto-laporan
	pv.GET("/read-path-foto", controllers.ReadFotoPembayaranvendor)

	//laporan-vendor
	//Create
	LV.POST("/input-lv", controllers.InputLaporanVendor)
	//read
	LV.GET("/Read-lv", controllers.ReadLaporanVendor)
	//update
	LV.PUT("/update-lv", controllers.UpdateLaporanVendor)
	//upload foto
	LV.POST("/upload-foto", controllers.UploadFotolaporanVendor)
	//Read-foto-laporan
	LV.GET("/read-path-foto", controllers.ReadFotolaporanVendor)
	//update_status_laporan_vendor
	LV.PUT("/update-status", controllers.UpdateStatusLaporanVendor)

	//Penawaran(header)
	//create
	ph.POST("/input-ph", controllers.InputHeaderPenawaran)
	//read
	ph.GET("/read-ph", controllers.ReadHeaderPenawaran)
	//update

	//Penawaran
	//create
	pen.POST("/input-pen", controllers.InputPenawaran)
	//read
	pen.GET("/read-pen", controllers.ReadPenawaran)
	//Update Status Penawaran
	pen.PUT("/update-status", controllers.UpdateStatusPenawaran)
	//update kop
	pen.PUT("/UpdateHeaderPenawaran", controllers.UpdateHeaderPenawaran)
	//update_judul
	pen.PUT("/update-judul", controllers.UpdateJudulPenawaran)
	//update item
	pen.PUT("/update-item", controllers.UpdateItemPenawaran)

	//Penjadwalan
	//input tanggal mulai
	PJDL.PUT("/tgl-ml", controllers.InputTanggalMulai)
	//read tanggal mulai
	PJDL.GET("/read-tgl-ml", controllers.ReadTanggalMulai)
	//read_judul_penawaran
	PJDL.GET("/read-judul-penawaran", controllers.ReadJudulPenawaran)
	//create task
	PJDL.POST("/input-task-penjadwalan", controllers.InputTaskPenjadwalan)
	//read task
	PJDL.GET("/read-task", controllers.ReadTask)
	//add dependentcies
	PJDL.PUT("/input-depedentcies", controllers.Inputdepedentcies)
	//generate jadwal cpm
	PJDL.GET("/generate-jadwal", controllers.GenerateJadwal)
	//Read_Jadwal
	PJDL.GET("/read-jadwal", controllers.ReadJadwal)
	//read_dep
	PJDL.GET("/read-dep", controllers.ReadDep)
	//edit

	//Laporan
	//create
	LP.POST("/input-lp", controllers.InputLaporan)
	//read
	LP.GET("/read-lp", controllers.ReadLaporan)
	//update
	LP.PUT("/update-lp", controllers.UpdateLaporan)
	//upload-foto-laporan
	LP.POST("upload-foto-laporan", controllers.UploadFotolaporan)
	//Read-foto-laporan
	LP.GET("/read-path-foto", controllers.ReadFotolaporan)
	//update_status_laporan
	LP.PUT("/update-status", controllers.UpdateStatusLaporan)

	//foto
	//get image foto
	FT.GET("/read-foto", controllers.ReadFoto)

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

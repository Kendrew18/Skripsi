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
	//update
	//delete
	//upload-foto-invoice
	pv.POST("/upload-fi", controllers.UploadInvoice)
	//read-foto-invoice
	pv.GET("/read-fi", controllers.Foto_Invoice)

	//laporan-vendor
	//Create
	LV.POST("/input-lv", controllers.InputLaporanVendor)
	//read
	LV.GET("/Read-lv", controllers.ReadLaporanVendor)
	//update
	LP.PUT("/update-lv", controllers.UpdateLaporanVendor)

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
	//update

	//Penjadwalan
	//create jadwal
	PJDL.POST("/input-task-penjadwalan", controllers.InputTaskPenjadwalan)
	//add dependentcies
	PJDL.PUT("/input-depedentcies", controllers.Inputdepedentcies)
	//edit

	//read

	//generate jadwal cpm
	PJDL.GET("/gene", controllers.GenerateJadwal)
	//input tanggal mulai
	PJDL.PUT("/tgl-ml", controllers.InputTanggalMulai)
	//read tanggal mulai
	PJDL.GET("/read-tgl-ml", controllers.ReadTanggalMulai)

	//Laporan
	//create
	LP.POST("/input-lp", controllers.InputLaporan)
	//read
	LP.GET("/read-lp", controllers.ReadLaporan)
	//update
	LP.PUT("/update-lp", controllers.UpdateLaporan)

	//Budgeting
	//input
	//update
	//read

	//kwitansi-pembayaran

	//PENDUKUNG//

	//Analisa Budgeting(perkiraan budget mingguan)
	//Read

	return e
}

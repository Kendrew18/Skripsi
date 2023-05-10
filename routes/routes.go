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

	//memberikan proges

	//input
	kv.POST("/input-kv", controllers.InputKontrakVendor)
	//read
	kv.GET("/read-kv", controllers.ReadKontrakVendor)

	//delete

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

	//Penawaran(header)
	//create
	ph.POST("/input-ph", controllers.InputHeaderPenawaran)
	//read
	ph.GET("/read-ph", controllers.ReadHeaderPenawaran)
	//update

	//Penawaran
	//create
	pen.POST("/input-pen", controllers.InputPenawaran)
	//update
	//read
	pen.GET("/read-pen", controllers.ReadPenawaran)

	//Penjadwalan
	//cpm

	//PENDUKUNG//

	//Laporan
	//read create update

	//Budgeting
	//input update read

	//Analisa Budgeting
	//Read

	return e
}

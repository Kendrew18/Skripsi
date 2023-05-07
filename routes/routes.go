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
	//update
	//delete
	//read
	kv.GET("/read-kv", controllers.ReadKontrakVendor)

	//Pemayaran-Vendor
	//input
	pv.POST("/input-pv", controllers.InputPembayaranVendor)
	//update
	//delete
	//read
	pv.GET("/read-pv", controllers.ReadPembayaranVendor)

	//Penawaran
	//create
	//update
	//read

	//Penjadwalan
	//cpm

	//PENDUKUNG//

	//Laporan
	// read create update

	//Budgeting
	//input update read

	//Analisa Budgeting
	//Read

	return e
}

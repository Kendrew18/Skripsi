package routes

import (
	controllers "Skripsi/controllers"
	"Skripsi/controllers/analisa_budgeting"
	"Skripsi/controllers/budgeting"
	"Skripsi/controllers/jadwal"
	"Skripsi/controllers/laporan_akhir"
	"Skripsi/controllers/penawaran"
	"Skripsi/controllers/proyek"
	"Skripsi/controllers/tagihan"
	"Skripsi/controllers/vendor_all"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-Skripsi-Project-Management")
	})

	kv := e.Group("/kv")
	pv := e.Group("/pv")
	LV := e.Group("/LV")
	FT := e.Group("/FT")
	MV := e.Group("/MV")
	BU := e.Group("/BU")
	TG := e.Group("/TG")
	AB := e.Group("/AB")
	LA := e.Group("/LA")

	//MAIN
	//User_Management(Ready)
	um := e.Group("/um")
	um.GET("/login-user", controllers.LoginUM)
	//Update_Token
	um.PUT("/update-token", controllers.UpdateToken)
	//Delete_Token
	um.PUT("/delete-token", controllers.DeleteToken)

	//Proyek(Ready)
	pryk := e.Group("/pryk")
	//Create
	pryk.POST("/input-proyek", proyek.InputProyek)
	//Read Nama Proyek
	pryk.GET("/read-nama", proyek.ReadNamaProyek)
	//Read Detail
	pryk.GET("/read-proyek", proyek.ReadProyek)
	//Finish (bag 2)
	pryk.PUT("/finish-proyek", proyek.FinishProyek)

	//History Proyek (bag 2)
	//Read Nama Proyek
	pryk.GET("/read-nama-his", proyek.ReadNamaProyekHistory)
	//Read detail history
	pryk.GET("/read-proyek-his", proyek.ReadHistory)

	//Edit (bag 2)

	//Kontrak-Vendor
	//Input_Kontrak_Vendor
	kv.POST("/input-kv", vendor_all.InputKontrakVendor)
	//Read_Kontrak_Vendor
	kv.GET("/read-kv", vendor_all.ReadKontrakVendor)
	//Delete_Kontrak_Vendor
	kv.DELETE("/delete-kontrak", vendor_all.DeleteKontrakVendor)
	//Pick_Vendor
	kv.GET("/pick-vendor", vendor_all.PickVendor)
	//Data_Filter
	kv.GET("/data-filter", vendor_all.DataFilter)

	//Pembayaran-Vendor
	//input
	pv.POST("/input-pv", vendor_all.InputPembayaranVendor)
	//read
	pv.GET("/read-pv", vendor_all.ReadPembayaranVendor)
	//Delete PEmbayaran Vendor
	pv.DELETE("/delete-pembayaran-vendor", vendor_all.DeletePembayaranVendor)
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
	LV.GET("/see-kontrak-vendor", vendor_all.SeeTaskVendor)
	//Delete Laporan Vendor
	LV.DELETE("/delete-laporan-vendor", vendor_all.DeleteLaporanVendor)

	//Penawaran (Ready)
	pen := e.Group("/pen")
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
	pen.POST("/input-tambahan-pekerjaan-tambah", penawaran.InputTambahanPekerjaanTambah)
	//pilih_judul_pekerjaan
	pen.GET("/pilih-judul-pekerjaan", penawaran.PilihJudulPekerjaan)

	//Penjadwalan (Ready)
	PJDL := e.Group("/PJDL")
	//Input-Durasi-Task
	PJDL.PUT("/input-durasi-task", jadwal.InputDurasitask)
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
	PJDL.PUT("/edit-dur-tgl", jadwal.EditDurTgl)
	//See-Calender-All
	PJDL.GET("/see-calender-all", jadwal.SeeCalenderAll)

	//Laporan
	LP := e.Group("/LP")
	//Input-Laporan
	LP.POST("/input-lp", jadwal.InputLaporan)
	//Read-Laporan
	LP.GET("/read-lp", jadwal.ReadLaporan)
	//Update-Laporan
	LP.POST("/update-lp", jadwal.UpdateLaporan)
	//upload-foto-laporan
	LP.POST("/upload-foto-laporan", jadwal.UploadFotolaporan)
	//Read-foto-laporan
	LP.GET("/read-path-foto", jadwal.ReadFotolaporan)
	//update_status_laporan
	LP.PUT("/update-status", jadwal.UpdateStatusLaporan)
	//See Task
	LP.GET("/see-task", jadwal.SeeTask)
	//Delete_Laporan
	LP.DELETE("/delete-laporan", jadwal.DeleteLaporan)

	//Master_Vendor
	//input-master-vendor
	MV.POST("/input-master-vendor", vendor_all.InputVendor)
	//read-master-vendor
	MV.GET("/read-master-vendor", vendor_all.ReadVendor)
	//delete-master-vendor
	MV.DELETE("/delete-master-vendor", vendor_all.DeleteVendor)
	//edit-master-vendor
	MV.PUT("/edit-master-vendor", vendor_all.EditVendor)

	//foto
	//get image foto
	FT.GET("/read-foto", vendor_all.ReadFoto)

	//Budgeting
	//Input_Realisasi
	BU.POST("/input-realisasi", budgeting.InputDetailBudgeting)
	//Read_Realisasi
	BU.GET("/read-realisasi", budgeting.ReadDetailBudgeting)
	//Read_Realisasi
	BU.DELETE("/delete-realisasi", budgeting.DeleteDetailBudgeting)
	//Read_Realisasi
	BU.PUT("/edit-realisasi", budgeting.UpdateDetailBudgeting)
	//Read_Budgeting
	BU.GET("/read-budgeting", budgeting.ReadBudgeting)
	//Pilih Kontra
	BU.GET("/pilih-kontrak", budgeting.PilihKontrak)

	//Tagihan
	//Input_Tagihan
	TG.POST("/input-tagihan", tagihan.InputTagihan)
	//Read_Realisasi
	TG.GET("/read-tagihan", tagihan.ReadTagihan)
	//Delete_Tagihan
	TG.DELETE("/delete-tagihan", tagihan.DeleteTagihan)
	//See_Judul
	TG.GET("/see-judul", tagihan.SeeJudul)
	//See_Sub_Pekerjaan
	TG.GET("/see-sub-pekerjaan", tagihan.SeeSubPekerjaan)

	//Analisa Budgeting(perkiraan budget mingguan)
	//Read
	AB.GET("/read-analisa-budgeting", analisa_budgeting.ReadAnalisaBudgeting)

	//Laporan_akhir
	//Read
	//Update
	LA.GET("/laporan-akhir", laporan_akhir.ReadLaporanAkhir)
	LA.PUT("/laporan-akhir", laporan_akhir.UpdateStatus)

	return e
}

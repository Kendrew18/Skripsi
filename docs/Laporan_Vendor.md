# How TO Use API Laporan
__________
##  Input Laporan Vendor

Link: kostsoda.onthewifi.com:38600/LV/input-lp

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")
	id_kontrak := c.FormValue("id_kontrak")
	check_Box := c.FormValue("check_Box")

##  Read Laporan Vendor

Link: kostsoda.onthewifi.com:38600/LV/Read-lv

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Update Laporan Vendor

Link: kostsoda.onthewifi.com:38600/LV/update-lv

Method: PUT

Controllers:

    id_laporan_vendor := c.FormValue("id_laporan_vendor")
    laporan := c.FormValue("laporan")
    id_kontrak := c.FormValue("id_kontrak")
	check_Box := c.FormValue("check_Box")

##  See Task Kontrak Vendor

Link: kostsoda.onthewifi.com:38600/LV/see-kontrak-vendor

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")
    id_laporan_vendor := c.FormValue("id_laporan_vendor")
    tanggal_laporan_vendor := c.FormValue("tanggal_laporan_vendor")

## Upload Foto

Link: kostsoda.onthewifi.com:38600/LV/upload-foto

Method: POST

Controllers:

    id_laporan_vendor := c.FormValue("id_laporan_vendor")
    photo := c.FormValue("photo")

##  Read Path Foto

Link: kostsoda.onthewifi.com:38600/LV/read-path-foto

Method: GET

Controllers:

    id_laporan_vendor := c.FormValue("id_laporan_vendor")

##  Delete Laporan Vendor

Link: kostsoda.onthewifi.com:38600/LV/delete-laporan-vendor

Method: DELETE

Controllers:

    id_laporan_vendor := c.FormValue("id_laporan_vendor")


# How TO Use API Laporan
__________
##  Input Durasi Pekerjaan

Link: kostsoda.onthewifi.com:38600/LP/input-lp

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	laporan := c.FormValue("laporan")
	tanggal_laporan := c.FormValue("tanggal_laporan")
	id_penjadwalan := c.FormValue("id_penjadwalan")
	check := c.FormValue("check")

##  Read Laporan

Link: kostsoda.onthewifi.com:38600/LP/read-lp

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Update Laporan

Link: kostsoda.onthewifi.com:38600/LP/update-lp

Method: PUT

Controllers:

    id_laporan := c.FormValue("id_laporan")
    laporan := c.FormValue("laporan")
    id_penjadwalan := c.FormValue("id_penjadwalan")
    check := c.FormValue("check")

##  See Task

Link: kostsoda.onthewifi.com:38600/LP/see-task

Method: GET

Controllers:

    tanggal_laporan := c.FormValue("tanggal_laporan")

##  Upload Foto Laporan

Link: kostsoda.onthewifi.com:38600/LP/upload-foto-laporan

Method: POST

Controllers:

    id_laporan := c.FormValue("id_laporan")
    photo := c.FormValue("photo")

## Delete Laporan

Link: kostsoda.onthewifi.com:38600/LP/delete-laporan

Method: DELETE

Controllers:

    id_laporan := c.FormValue("id_laporan")


    
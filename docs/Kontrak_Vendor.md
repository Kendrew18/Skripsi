# How TO Use API Kontrak Vendor
__________
##  Input Kontrak Vendor

Link: kostsoda.onthewifi.com:38600/kv/input-kv

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_master_vendor := c.FormValue("id_master_vendor")
	total_nilai_kontrak := c.FormValue("total_nilai_kontrak")
	tanggal_mulai_kontrak := c.FormValue("tanggal_mulai_kontrak")
	tanggal_berakhir_kontrak := c.FormValue("tanggal_berakhir_kontrak")
	tanggal_pengiriman := c.FormValue("tanggal_pengiriman")
	tanggal_pengerjaan_dimulai := c.FormValue("tanggal_pengerjaan_dimulai")
	tanggal_pengerjaan_selesai := c.FormValue("tanggal_pengerjaan_selesai")

##  Read Kontrak Vendor

Link: kostsoda.onthewifi.com:38600/kv/read-kv

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Delete Kontrak Vendor

Link: kostsoda.onthewifi.com:38600/kv/delete-kontrak

Method: DELETE

Controllers:

    id_kontrak := c.FormValue("id_kontrak")

##  Pick Vendor

Link: kostsoda.onthewifi.com:38600/kv/pick-vendor

Method: GET

Controllers:-




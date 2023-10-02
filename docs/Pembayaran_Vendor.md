# How TO Use API Pembayaran Vendor
__________
##  Input Pembayaran Vendor

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
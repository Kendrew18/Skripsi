# How TO Use API Budgeting
__________
##  Input Budgeting
Link: kostsoda.onthewifi.com:38600/BU/input-realisasi

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

##  Read Realisasi
Link: kostsoda.onthewifi.com:38600/BU/read-realisasi

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

##  Delete Realisasi
Link: kostsoda.onthewifi.com:38600/BU/delete-realisasi

Method: Delete

Controllers:

    id_realisasi := c.FormValue("id_realisasi")

##  Update Realisasi
Link: kostsoda.onthewifi.com:38600/BU/delete-realisasi

Method: PUT

Controllers:

    id_realisasi := c.FormValue("id_realisasi")
	id_kontrak := c.FormValue("id_kontrak")
	perihal_pengeluaran := c.FormValue("perihal_pengeluaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_pembayaran := c.FormValue("nominal_pembayaran")
	catatan := c.FormValue("catatan")

##  Read Budgeting
Link: kostsoda.onthewifi.com:38600/BU/read-budgeting

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Pilih Kontrak
Link: kostsoda.onthewifi.com:38600/BU/pilih-kontrak

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")
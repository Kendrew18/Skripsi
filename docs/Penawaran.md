# How TO Use API Penawaran
__________
##  Input Penawaran

Link: kostsoda.onthewifi.com:38600/pen/input-pen

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	judul := c.FormValue("judul")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah") (double)
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total") (double)

NB: sub_pekerjaan, catatan, jumlah, satuan, sub_total berupa String Separator

##  Input Sub Penawaran

Link: kostsoda.onthewifi.com:38600/pen/input-sub-pen

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah") (double)
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total") (double)

##  Read Penawaran

Link: kostsoda.onthewifi.com:38600/pen/read-pen

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Pilih Judul Pekerjaan

Link: kostsoda.onthewifi.com:38600/pen/pilih-judul-pekerjaan

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Update Status

Link: kostsoda.onthewifi.com:38600/pen/update-status

Method: PUT

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Update Judul

Link: kostsoda.onthewifi.com:38600/pen/update-judul

Method: PUT

Controllers:

    id_proyek := c.FormValue("id_proyek")
    judul := c.FormValue("judul")

##  Update Item

Link: kostsoda.onthewifi.com:38600/pen/update-item

Method: PUT

Controllers:

    id_penawaran := c.FormValue("id_penawaran")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah") (double)
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")(double)

##  Input Tambahan Sub Pekerjaan

Link: kostsoda.onthewifi.com:38600/pen/input-tambahan-sub-pekerjaan

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah") (double)
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total") (double)
	tanggal_pekerjaan_dimulai := c.FormValue("tanggal_pekerjaan_dimulai")
	durasi := c.FormValue("durasi")

##  Input Tambahan Pekerjaan Tambah

Link: kostsoda.onthewifi.com:38600/pen/input-tambahan-pekerjaan-tambah

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	id_penawaran := c.FormValue("id_penawaran")
	sub_pekerjaan := c.FormValue("sub_pekerjaan")
	catatan := c.FormValue("catatan")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	harga := c.FormValue("harga")
	sub_total := c.FormValue("sub_total")
	tanggal_pekerjaan_dimulai := c.FormValue("tanggal_pekerjaan_dimulai")
	durasi := c.FormValue("durasi")

NB: sub_pekerjaan, catatan, jumlah, satuan, sub_total, tanggal_pekerjaan_dimulai, durasi berupa String Separator

##  Pilih Judul Pekerjaan

Link: kostsoda.onthewifi.com:38600/pen/pilih-judul-pekerjaan

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")
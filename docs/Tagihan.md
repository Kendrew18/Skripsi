# How TO Use API Tagihan
__________
##  Input Tagihan

Link: kostsoda.onthewifi.com:38600/TG/input-tagihan

Method: POST

Controllers:

    id_proyek := c.FormValue("id_proyek")
	perihal := c.FormValue("perihal")
	tanggal_pemberian_kwitansi := c.FormValue("tanggal_pemberian_kwitansi")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	nominal_keseluruhan := c.FormValue("nominal_keseluruhan")
	id_penawaran := c.FormValue("id_penawaran")
	id_sub_pekerjaan := c.FormValue("id_sub_pekerjaan")
	nominal := c.FormValue("nominal")

NB: id_penawaran, id_sub_pekerjaan, nominal berupa String Separator

##  Read Tagihan

Link: kostsoda.onthewifi.com:38600/TG/read-tagihan

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Delete Tagihan

Link: kostsoda.onthewifi.com:38600/TG/delete-tagihan

Method: Delete

Controllers:

    id_tagihan := c.FormValue("id_tagihan")

##  See Judul

Link: kostsoda.onthewifi.com:38600/TG/see-judul

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  See Sub Pekerjaan

Link: kostsoda.onthewifi.com:38600/TG/see-sub-pekerjaan

Method: GET

Controllers:

    id_penawaran := c.FormValue("id_penawaran")

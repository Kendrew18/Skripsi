# How TO Use API Proyek
__________
##  Input Proyek

Link: kostsoda.onthewifi.com:3333/pryk/input-proyek

Method: POST

Controllers:

    id_user := c.FormValue("id_user")
	nama_proyek := c.FormValue("nama_proyek")
	jumlah_lantai := c.FormValue("jumlah_lantai")
	luas_tanah := c.FormValue("luas_tanah")
	nama_penanggungjawab_proyek := c.FormValue("nama_penanggungjawab")
	nama_client := c.FormValue("nama_client")
	jenis_gedung := c.FormValue("jenis_gedung")
	alamat := c.FormValue("alamat")
	tanggal_mulai_kerja := c.FormValue("tanggal_mulai_kerja")

##  Read Nama Proyek

Link: kostsoda.onthewifi.com:3333/pryk/read-nama

Method: GET

Controllers:

##  Read Detail Proyek

Link: kostsoda.onthewifi.com:3333/pryk/read-proyek

Method: GET

Controllers:

    id_proyek := c.FormValue("id_proyek")

##  Upadate Status Proyek

Link: kostsoda.onthewifi.com:3333/pryk/finish-proyek

Method: PUT

Controllers:

    id_proyek := c.FormValue("id_proyek")
	

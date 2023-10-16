# How TO Use API Pembayaran Vendor
__________
##  Input Pembayaran Vendor

Link: kostsoda.onthewifi.com:38600/pv/input-pv

Method: POST

Controllers:

    id_kontrak := c.FormValue("id_kontrak")
	nomor_invoice := c.FormValue("nomor_invoice")
	jumlah_pembayaran := c.FormValue("jumlah_pembayaran")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")

##  Read Pembayaran Vendor

Link: kostsoda.onthewifi.com:38600/pv/read-pv

Method: GET

Controllers:

    id_kontrak := c.FormValue("id_kontrak")

##  Upload Bukti Pembayaran

Link: kostsoda.onthewifi.com:38600/pv/read-pv

Method: POST

Controllers:

    id_PV := c.FormValue("id_PV")

##  Read Path Foto

Link: kostsoda.onthewifi.com:38600/pv/read-path-foto

Method: GET

Controllers:

    id_PV := c.FormValue("id_PV")

##  Delete Pembayaran Vendor

Link: kostsoda.onthewifi.com:38600/pv/delete-pembayaran-vendor

Method: DELETE

Controllers:

    id_PV := c.FormValue("id_PV")
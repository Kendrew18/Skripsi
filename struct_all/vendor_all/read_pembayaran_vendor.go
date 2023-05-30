package vendor_all

type Read_Pembayaran_Vendor struct {
	Id_PV              string `json:"id_pv"`
	Nomor_invoice      string `json:"nomor_invoice"`
	Jumlah_Pembayaran  int64  `json:"jumlah_pembayaran"`
	Tanggal_Pembayaran string `json:"tanggal_pembayaran"`
	Foto_invoice       string `json:"foto_invoice"`
}

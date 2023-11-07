package vendor_all

type Read_Kontrak_Vendor struct {
	Id_kontak                string                `json:"id_kontak"`
	Nama_vendor              string                `json:"nama_vendor"`
	Pekerjaan_Vendor         string                `json:"pekerjaan_vendor"`
	Tanggal_mulai_kontrak    string                `json:"tanggal_mulai_kontrak"`
	Tanggal_berakhir_kontrak string                `json:"tanggal_berakhir_kontrak"`
	Detail_Kontrak_Vendor    Detail_Kontrak_Vendor `json:"detail_kontrak_vendor"`
}

type Detail_Kontrak_Vendor struct {
	Total_nilai_kontrak         int64  `json:"total_nilai_kontrak"`
	Nomial_Pembayaran           int64  `json:"nomial_pembayaran"`
	Tanggal_Pengiriman          string `json:"jenis_pekerjaan"`
	Tanggal_mulai_pengerjaan    string `json:"tanggal_mulai_kontrak"`
	Tanggal_berakhir_pengerjaan string `json:"tanggal_berakhir_kontrak"`
	Sisa_pembayaran             int64  `json:"sisa_pembayaran"`
}

type Filter_Kontrak struct {
	Id_master_vendor string `json:"id_master_vendor"`
	Nama_vendor      string `json:"nama_vendor"`
}

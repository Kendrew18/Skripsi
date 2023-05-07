package struct_all

type Read_Kontrak_Vendor struct {
	Id_proyek                string `json:"id_proyek"`
	Id_kontak                string `json:"id_kontak"`
	Nomor_Kontrak            string `json:"nomor_kontrak"`
	Nama_vendor              string `json:"nama_vendor"`
	Total_nilai_kontrak      int64  `json:"total_nilai_kontrak"`
	Nomial_Pembayaran        int64  `json:"nomial_pembayaran"`
	Jenis_Pekerjaan          string `json:"jenis_pekerjaan"`
	Tanggal_mulai_kontrak    string `json:"tanggal_mulai_kontrak"`
	Tanggal_berakhir_kontrak string `json:"tanggal_berakhir_kontrak"`
	Sisa_pembayaran          int64  `json:"sisa_pembayaran"`
}

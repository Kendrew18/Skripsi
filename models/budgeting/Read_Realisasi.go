package budgeting

type Read_Realisasi struct {
	Id_Realisasi        string `json:"id_realisasi"`
	Id_Proyek           string `json:"id_proyek"`
	Id_Sub_Pekerjaan    string `json:"id_sub_pekerjaan"`
	Id_Kontrak          string `json:"id_kontrak"`
	Nama_Vendor         string `json:"nama_vendor"`
	Perihal_Pengeluaran string `json:"perihal_pengeluaran"`
	Nominal_Pembayaran  int64  `json:"nominal_pembayaran"`
	Catatan             string `json:"catatan"`
}

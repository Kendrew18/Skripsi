package budgeting

type Read_Sub_Pekerjaan struct {
	Id_sub_pekerjaan string `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string `json:"sub_pekerjaan"`
}

type Read_Pengeluaran struct {
	Id_laporan         string               `json:"id_laporan"`
	Tanggal            string               `json:"tanggal"`
	Read_Sub_Pekerjaan []Read_Sub_Pekerjaan `json:"read_sub_pekerjaan"`
}

type Kontrak_Vendor struct {
	Id_kontak_vendor string `json:"id_kontak_vendor"`
	Nama_vendor      string `json:"nama_vendor"`
}

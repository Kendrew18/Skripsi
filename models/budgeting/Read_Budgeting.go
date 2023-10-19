package budgeting

type Read_Sub_Pekerjaan struct {
	Id_sub_pekerjaan string `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string `json:"sub_pekerjaan"`
	Biaya_Estimasi   int64  `json:"biaya_estimasi"`
	Biaya_Realisasi  int64  `json:"biaya_realisasi"`
	Biaya_Pelunasan  int64  `json:"biaya_pelunasan"`
}

type Read_Budgeting struct {
	Id_penawaran       string               `json:"id_penawaran"`
	Judul              string               `json:"judul"`
	Read_Sub_Pekerjaan []Read_Sub_Pekerjaan `json:"read_sub_pekerjaan"`
}

type Kontrak_Vendor struct {
	Id_kontak_vendor string `json:"id_kontak_vendor"`
	Nama_vendor      string `json:"nama_vendor"`
}

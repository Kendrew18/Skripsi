package penawaran

type Read_Penawaran_String struct {
	Id_sub_pekerjaan string `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string `json:"sub_pekerjaan"`
	Keterangan       string `json:"keterangan"`
	Jumlah           string `json:"jumlah"`
	Satuan           string `json:"satuan"`
	Harga            string `json:"harga"`
	Total            string `json:"total"`
}

type Read_Penawaran struct {
	Id_penawaran     string    `json:"id_penawaran"`
	Judul            string    `json:"judul"`
	Id_sub_pekerjaan []string  `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    []string  `json:"sub_pekerjaan"`
	Keterangan       []string  `json:"keterangan"`
	Jumlah           []float64 `json:"jumlah"`
	Satuan           []string  `json:"satuan"`
	Harga            []float64 `json:"harga"`
	Total            []float64 `json:"total"`
	Sub_total        int64     `json:"sub_total"`
}

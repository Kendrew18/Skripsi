package penawaran

type Read_Penawaran_String struct {
	Id_sub_pekerjaan string `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string `json:"sub_pekerjaan"`
	Keterangan       string `json:"keterangan"`
	Jumlah           string `json:"jumlah"`
	Satuan           string `json:"satuan"`
	Harga            string `json:"harga"`
	Sub_total        string `json:"total"`
}

type Read_Penawaran struct {
	Id_penawaran     string    `json:"id_penawaran"`
	Judul            string    `json:"judul"`
	Id_sub_pekerjaan []string  `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    []string  `json:"sub_pekerjaan"`
	Catatan          []string  `json:"catatan"`
	Jumlah           []float64 `json:"jumlah"`
	Satuan           []string  `json:"satuan"`
	Harga            []int     `json:"harga"`
	Sub_total        []int64   `json:"total"`
	Total            int64     `json:"sub_total"`
}

package struct_all

type Read_Penawaran_String struct {
	Sub_pekerjaan string `json:"sub_pekerjaan"`
	Keterangan    string `json:"keterangan"`
	Jumlah        string `json:"jumlah"`
	Satuan        string `json:"satuan"`
	Harga         string `json:"harga"`
	Total         string `json:"total"`
}

type Read_Penawaran struct {
	Judul         string    `json:"judul"`
	Sub_pekerjaan []string  `json:"sub_pekerjaan"`
	Keterangan    []string  `json:"keterangan"`
	Jumlah        []float64 `json:"jumlah"`
	Satuan        []string  `json:"satuan"`
	Harga         []int64   `json:"harga"`
	Total         []int64   `json:"total"`
	Sub_total     int64     `json:"sub_total"`
}

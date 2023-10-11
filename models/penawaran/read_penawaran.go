package penawaran

type Read_Detail_Penawaran struct {
	Id_sub_pekerjaan string  `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string  `json:"sub_pekerjaan"`
	Catatan          string  `json:"Catatan"`
	Jumlah           float64 `json:"jumlah"`
	Satuan           string  `json:"satuan"`
	Harga            int64   `json:"harga"`
	Sub_total        int64   `json:"Sub_total"`
	Status           int     `json:"status"`
}

type Read_Penawaran struct {
	Id_penawaran          string                  `json:"id_penawaran"`
	Judul                 string                  `json:"judul"`
	Read_Detail_Penawaran []Read_Detail_Penawaran `json:"read_detail_penawaran"`
	Total                 int64                   `json:"total"`
	Status                int                     `json:"status"`
}

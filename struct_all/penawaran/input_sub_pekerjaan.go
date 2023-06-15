package penawaran

type Read_Penawaran_Input_Sub_Pekerjaan struct {
	Id_sub_pekerjaan string `json:"id_sub_pekerjaan"`
	Sub_pekerjaan    string `json:"sub_pekerjaan"`
	Catatan          string `json:"catatan"`
	Jumlah           string `json:"jumlah"`
	Satuan           string `json:"satuan"`
	Harga            string `json:"harga"`
	Sub_total        string `json:"total"`
	Total            int64  `json:"sub_total"`
}

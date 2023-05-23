package struct_all

type Read_proyek struct {
	Id_proyek        string  `json:"id_proyek"`
	Nama_proyek      string  `json:"nama_proyek"`
	Jumlah_lantai    int     `json:"jumlah_lantai"`
	Luas_tanah       float64 `json:"luas_tanah"`
	Penangung_Jawab  string  `json:"penangung_jawab"`
	Status_penawaran int     `json:"status_penawaran"`
}

package proyek

type Read_proyek struct {
	Id_proyek              string  `json:"id_proyek"`
	Nama_proyek            string  `json:"nama_proyek"`
	Nama_Client_Perusahaan string  `json:"nama_client_perusahaan"`
	Jenis_gedung           string  `json:"jenis_gedung"`
	Alamat                 string  `json:"alamat"`
	Luas_tanah             float64 `json:"luas_tanah"`
	Jumlah_lantai          int     `json:"jumlah_lantai"`
	Penangung_Jawab        string  `json:"penangung_jawab"`
	Tanggal_mulai_kerja    string  `json:"tanggal_mulai_kerja"`
	Status_penawaran       int     `json:"status_penawaran"`
}

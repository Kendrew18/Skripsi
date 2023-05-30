package jadwal

type Read_Jadwal struct {
	Id_penjadwalan   string `json:"id_penjadwalan"`
	Nama_Task        string `json:"nama_task"`
	Durasi_Task      string `json:"durasi_task"`
	Tanggal_mulai    string `json:"tanggal_mulai"`
	Tanggal_berakhir string `json:"tanggal_berakhir"`
}

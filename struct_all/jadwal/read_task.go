package jadwal

type Sub_Task struct {
	Id_penjadwalan string `json:"id_penjadwalan"`
	Nama_Task      string `json:"nama_task"`
	Durasi_Task    string `json:"durasi_task"`
	Dependentcies  string `json:"dependentcies"`
}

type Read_Task struct {
	Id_penawaran    string     `json:"id_penawaran"`
	Judul_penawaran string     `json:"judul_penawaran"`
	Sub_task        []Sub_Task `json:"sub_task"`
}

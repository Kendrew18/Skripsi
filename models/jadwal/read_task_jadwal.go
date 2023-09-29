package jadwal

type Sub_Task_Jadwal struct {
	Id_penjadwalan  string `json:"id_penjadwalan"`
	Nama_Task       string `json:"nama_task"`
	Tanggal_Mulai   string `json:"Tanggal_Mulai"`
	Tanggal_Selesai string `json:"Tanggal_Selesai"`
}

type Read_Task_Jadwal struct {
	Id_penawaran    string            `json:"id_penawaran"`
	Judul_penawaran string            `json:"judul_penawaran"`
	Sub_task        []Sub_Task_Jadwal `json:"sub_task"`
}

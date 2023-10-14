package jadwal

type Read_Task_Laporan struct {
	Id_penjadwalan string `json:"id_penjadwalan"`
	Nama_Task      string `json:"nama_task"`
	Progress       int    `json:"progress"`
	Check_box      int    `json:"check_box"`
}

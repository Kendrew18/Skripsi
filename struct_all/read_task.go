package struct_all

type Read_Task struct {
	Id_penjadwalan string `json:"id_penjadwalan"`
	Nama_Task      string `json:"nama_task"`
	Durasi_Task    string `json:"durasi_task"`
	Dependentcies  string `json:"dependentcies"`
}

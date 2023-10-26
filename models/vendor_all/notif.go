package vendor_all

type Read_All_Notif struct {
	Tanggal    string       `json:"tanggal"`
	Read_Notif []Read_Notif `json:"read_notif"`
}

type Read_Notif struct {
	Id_notif   string `json:"id_notif"`
	Id_kontrak string `json:"id_kontrak"`
	Tanggal    string `json:"tanggal"`
	Pesan      string `json:"pesan"`
	Jam_Notif  string `json:"jam_notif"`
}

type Read_Notif_Pop_up struct {
	Id_notif string `json:"id_notif"`
	Tanggal  string `json:"tanggal"`
	Pesan    string `json:"pesan"`
	Status_1 int    `json:"status_1"`
	Status_2 int    `json:"status_2"`
}

package vendor_all

type Read_Task_Laporan_Vendor struct {
	Id_Kontrak       string `json:"id_kontrak"`
	Nama_Vendor      string `json:"nama_vendor"`
	Pekerjaan_Vendor string `json:"pekerjaan_vendor"`
	Progress         int    `json:"progress"`
	Check_Box        int    `json:"check_box"`
}

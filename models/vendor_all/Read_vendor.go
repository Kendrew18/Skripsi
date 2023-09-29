package vendor_all

type Read_Vendor struct {
	Id_Master_Vendor   string                   `json:"id_master_vendor"`
	Nama_Vendor        string                   `json:"nama_vendor"`
	Pekerjaan_Vendor   string                   `json:"pekerjaan_vendor"`
	Pekerjaan_selesai  int                      `json:"pekerjaan_selesai"`
	Pekerjaan_berjalan int                      `json:"pekerjaan_berjalan"`
	Detail_Vendor      []Detail_Read_Vendor_Fix `json:"detail_read_vendor"`
}

type Detail_Read_Vendor struct {
	Id_proyek   string `json:"Id_proyek"`
	Nama_proyek string `json:"Nama_proyek"`
	Progres     string `json:"progres"`
}

type Detail_Read_Vendor_Fix struct {
	Id_proyek   string `json:"Id_proyek"`
	Nama_proyek string `json:"Nama_proyek"`
	Progres     int    `json:"progres"`
}

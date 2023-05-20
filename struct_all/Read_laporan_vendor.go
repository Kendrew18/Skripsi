package struct_all

type Read_Laporan_Vendor struct {
	Id_laporan_vendor string `json:"id_laporan_vendor"`
	Nama_vendor       string `json:"nama_vendor"`
	Pekerjaan_vendor  string `json:"pekerjaan_vendor"`
	Laporan           string `json:"laporan"`
	Tanggal_laporan   string `json:"tanggal_laporan"`
	Status_laporan    int    `json:"status_laporan"`
}

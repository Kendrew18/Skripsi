package jadwal

type Read_Laporan struct {
	Id_laporan      string `json:"id_laporan"`
	Laporan         string `json:"laporan"`
	Tanggal_laporan string `json:"tanggal_laporan"`
	Status_laporan  int    `json:"status_laporan"`
}
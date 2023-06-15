package jadwal

type Read_Laporan_String struct {
	Id_Penjadwalan string `json:"id_penjadwalan"`
}

type Read_Laporan struct {
	Id_laporan      string           `json:"id_laporan"`
	Laporan         string           `json:"laporan"`
	Tanggal_laporan string           `json:"tanggal_laporan"`
	Status_laporan  int              `json:"status_laporan"`
	Detail_Laporan  []Detail_Laporan `json:"detail_laporan"`
}

type Detail_Laporan struct {
	Id_Penjadwalan     string `json:"id_penjadwalan"`
	Nama_Sub_Pekerjaan string `json:"nama_sub_pekerjaan"`
	Progress           int    `json:"progress"`
}

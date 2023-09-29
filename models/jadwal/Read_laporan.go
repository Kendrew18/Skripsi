package jadwal

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

type Detail_Laporan_TB struct {
	Id_Penjadwalan string `json:"id_penjadwalan"`
}
type Detail_Laporan_Update struct {
	Co                int    `json:"co"`
	Id_detail_laporan string `json:"id_detail_laporan"`
	Id_Penjadwalan    string `json:"id_penjadwalan"`
	Check_Box         int    `json:"check_box"`
}

type Id_Detail_Laporan struct {
	Id_detail_laporan string `json:"id_detail_laporan"`
}

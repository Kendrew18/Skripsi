package budgeting

type Analisa_Budgeting struct {
	Tanggal_Awal   string `json:"tanggal_mulai"`
	Tanggal_Akhir  string `json:"tanggal_akhir"`
	Biaya_Mingguan int64  `json:"biaya_mingguan"`
	CV             int64  `json:"cv"`
	SV             int64  `json:"sv"`
	CPI            int64  `json:"cpi"`
	SPI            int64  `json:"spi"`
}

type Detail_Analisa_Budgeting struct {
	Id_penawaran         string                 `json:"id_penawaran"`
	Nama_judul           string                 `json:"nama_judul"`
	Detail_Sub_Pekerjaan []Detail_Sub_Pekerjaan `json:"detail_sub_pekerjaan"`
}

type Detail_Sub_Pekerjaan struct {
	Id_Sub_Pekerjaan   string `json:"id_sub_pekerjaan"`
	Nama_Sub_Pekerjaan string `json:"nama_sub_pekerjaan"`
	Progress           string `json:"progress"`
	PV                 int64  `json:"pv"`
	EV                 int64  `json:"ev"`
	AC                 int64  `json:"ac"`
}

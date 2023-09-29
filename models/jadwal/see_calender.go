package jadwal

type See_Calender struct {
	Judul           string `json:"judul"`
	Tanggal_mulai   string `json:"tanggal_mulai"`
	Tanggal_selesai string `json:"tanggal_selesai"`
}

type See_Calender_Vendor struct {
	Id_kontrak         string `json:"Id_kontrak"`
	Nama_vendor        string `json:"Nama_vendor"`
	Nominal_Pembayaran string `json:"Nominal_Pembayaran"`
	Tanggal_mulai      string `json:"tanggal_mulai"`
	Tanggal_selesai    string `json:"tanggal_selesai"`
}

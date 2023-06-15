package vendor_all

type Read_Laporan_Vendor struct {
	Id_laporan_vendor     string                  `json:"id_laporan_vendor"`
	Laporan               string                  `json:"laporan"`
	Tanggal_laporan       string                  `json:"tanggal_laporan"`
	Status_laporan        int                     `json:"status_laporan"`
	Detail_Laporan_Vendor []Detail_Laporan_Vendor `json:"detail_laporan_vendor"`
}

type Detail_Laporan_Vendor struct {
	Id_kontrak       string `json:"id_kontrak"`
	Nama_vendor      string `json:"nama_vendor"`
	Pekerjaan_vendor string `json:"pekerjaan_vendor"`
	Progress         int    `json:"progress"`
}

type Read_Laporan_Vendor_String struct {
	Id_Kontrak_Vendor string `json:"Id_Kontrak_Vendor"`
}

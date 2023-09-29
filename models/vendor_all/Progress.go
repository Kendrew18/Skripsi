package vendor_all

type Progress_Vendor struct {
	Id_kontrak       string `json:"id_kontrak"`
	Nama_Vendor      string `json:"nama_vendor"`
	Working_Progess  int    `json:"working_progess"`
	Working_Complate int    `json:"working_complate"`
	Durasi           int    `json:"durasi"`
}

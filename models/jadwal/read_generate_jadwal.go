package jadwal

type Gene_JDL struct {
	Id               string `json:"id"`
	Durasi           int    `json:"durasi"`
	Dependentcies    string `json:"dependentcies"`
	Status_urutan    int    `json:"status_urutan"`
	Es               int    `json:"es"`
	Ls               int    `json:"ls"`
	Ef               int    `json:"ef"`
	Lf               int    `json:"lf"`
	Tf               int    `json:"tf"`
	Ff               int    `json:"ff"`
	Tanggal_mulai    string `json:"tanggal_mulai"`
	Tanggal_berakhir string `json:"tanggal_berakhir"`
	CPM              int    `json:"CPM"`
}

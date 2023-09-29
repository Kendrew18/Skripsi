package jadwal

type Progress struct {
	Id_penjadwalan string `json:"id_penjadwalan"`
	Progress       int    `json:"progress"`
	Durasi         int    `json:"durasi"`
	Complate       int    `json:"complate"`
}

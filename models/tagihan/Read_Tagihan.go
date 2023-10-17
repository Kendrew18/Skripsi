package tagihan

type Read_Tagihan struct {
	Id_Tagihan                 string                `json:"id_tagihan"`
	Perihal_Tagihan            string                `json:"perihal_tagihan"`
	Tanggal_Pemberian_Kwitansi string                `json:"tanggal_pemberian_kwitansi"`
	Tanggal_Pembayaran         string                `json:"tanggal_pembayaran"`
	Nominal_Keseluruhan        int64                 `json:"nominal_keseluruhan"`
	Read_Detail_Tagihan        []Read_Detail_Tagihan `json:"read_detail_tagihan"`
}

type Read_Detail_Tagihan struct {
	Id_detail_tagihan  string `json:"id_detail_tagihan"`
	Id_Penawaran       string `json:"id_penawaran"`
	Id_Sub_Pekerjaan   string `json:"id_sub_pekerjaan"`
	Nama_Sub_Pekerjaan string `json:"nama_sub_pekerjaan"`
	Nominal            int64  `json:"nominal"`
}

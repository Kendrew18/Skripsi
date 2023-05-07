package struct_all

type Read_Header struct {
	Id_Header_Penawwaran string `json:"id_header_penawwaran"`
	Kode_surat           string `json:"kode_surat"`
	Tanggal_dibuat       string `json:"tanggal_dibuat"`
	Nama_Perusahaan      string `json:"nama_perusahaan"`
	Alamat_Perusahaan    string `json:"alamat_perusahaan"`
}

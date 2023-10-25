package models

type Id_user struct {
	Id_user     string `json:"id_user"`
	Status_akun int    `json:"status_akun"`
}

type User struct {
	Id_user     string `json:"id_user"`
	Nama_user   string `json:"nama_user"`
	Token_user  string `json:"token_user"`
	Status_akun int    `json:"status_akun"`
}

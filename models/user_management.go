package models

import (
	"Skripsi/db"
	"Skripsi/struct_all"
	"Skripsi/tools"
	"fmt"
	"net/http"
)

func Login(username string, password string) (tools.Response, error) {
	var arr struct_all.Id_user
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT kode_user,status FROM user where username=? && password_user=?"

	err := con.QueryRow(sqlStatement, username, password).Scan(&arr.Id_user, &arr.Status_akun)

	if err != nil || arr.Id_user == "" {
		arr.Id_user = ""
		arr.Status_akun = 0
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	fmt.Println(arr)
	return res, nil
}

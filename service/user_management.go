package service

import (
	"Skripsi/config/db"
	"Skripsi/models"
	"Skripsi/service/tools"
	"fmt"
	"net/http"
)

//Login
func Login(username string, password string) (tools.Response, error) {
	var arr models.Id_user
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT kode_user,status_user FROM user where username=? && password_user=?"

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

//Update Token
func Update_Token(id_user string, token string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE user SET token=? WHERE kode_user=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(token, id_user)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}

//Delete Token
func Delete_Token(id_user string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE user SET token=? WHERE kode_user=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec("", id_user)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}

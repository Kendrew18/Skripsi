package vendor_all

import (
	"Skripsi/config/db"
	str "Skripsi/models"
	"Skripsi/models/vendor_all"
	"Skripsi/service/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Input Pembayaran Vendor
func Input_Pembayaran_Vendor(id_kontrak string, nomor_invoice string, jumlah_pembayaran int64, tanggal_pembayaran string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM pembayaran_vendor ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_PV := "PV-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO pembayaran_vendor (co,id_PV,id_kontrak,nomor_invoice,jumlah_pembayaran,tanggal_pembayaran ,foto_invoice) values(?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_PV, id_kontrak, nomor_invoice, jumlah_pembayaran, date_sql, "")

	if err != nil {
		return res, err
	}

	Sisa := int64(0)

	sqlStatement = "SELECT sisa_pembayaran FROM kontrak_vendor WHERE id_kontrak=?"
	err = con.QueryRow(sqlStatement, id_kontrak).Scan(&Sisa)
	if err != nil {
		return res, err
	}

	Sisa_sekarang := Sisa - jumlah_pembayaran

	sqlStatement = "UPDATE kontrak_vendor SET sisa_pembayaran=? WHERE id_kontrak=?"
	stmt, err = con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	_, err = stmt.Exec(Sisa_sekarang, id_kontrak)
	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read Pembayaran Vendor
func Read_Pembayaran_Vendor(id_kontrak string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Read_Pembayaran_Vendor
	var invent vendor_all.Read_Pembayaran_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_PV, nomor_invoice, jumlah_pembayaran, DATE_FORMAT(tanggal_pembayaran, '%d-%m%-%Y') FROM pembayaran_vendor WHERE id_kontrak=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_kontrak)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_PV, &invent.Nomor_invoice, &invent.Jumlah_Pembayaran, &invent.Tanggal_Pembayaran)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}

//Upload Vendor
func Upload_Invoice(id_PV string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()
	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""

	if strings.Contains(handler.Filename, "jpg") {
		path = "uploads/foto_pembayaran_vendor/" + id_PV + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/foto_pembayaran_vendor/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/foto_pembayaran_vendor/" + id_PV + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/foto_pembayaran_vendor/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/foto_pembayaran_vendor/" + id_PV + ".png"
		tempFile, err = ioutil.TempFile("uploads/foto_pembayaran_vendor/", "Read"+"*.png")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println("new path:", tempFile.Name())

	sqlstatement := "UPDATE pembayaran_vendor SET foto_invoice=? WHERE id_PV=?"
	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(path, id_PV)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read Foto Pembayarn Vendor
func Read_Foto_Pembayaran_vendor(id_pembayaran_vendor string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Foto
	var invent str.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT foto_invoice FROM pembayaran_vendor WHERE id_PV=? "

	err := con.QueryRow(sqlStatement, id_pembayaran_vendor).Scan(&invent.Path_foto)

	if err != nil {
		return res, err
	}

	arr_invent = append(arr_invent, invent)

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}

//Delete Pembayaran Vendor
func Delete_Pembayaran_Vendor(id_pembayaran_vendor string) (tools.Response, error) {
	var res tools.Response
	var obj vendor_all.Read_Pembayaran_Vendor

	con := db.CreateCon()

	ID_Kontrak := ""

	sqlStatement := "SELECT id_PV,jumlah_pembayaran,id_kontrak FROM pembayaran_vendor WHERE id_PV=? "

	err := con.QueryRow(sqlStatement, id_pembayaran_vendor).Scan(&obj.Id_PV,
		&obj.Jumlah_Pembayaran, &ID_Kontrak)

	if err != nil {
		return res, err
	}

	temp_sisa_pembayaran := int64(0)

	sqlStatement = "SELECT sisa_pembayaran FROM kontrak_vendor WHERE id_kontrak=? "

	err = con.QueryRow(sqlStatement, ID_Kontrak).Scan(&temp_sisa_pembayaran)

	temp_sisa_pembayaran = temp_sisa_pembayaran + obj.Jumlah_Pembayaran

	sqlstatement := "UPDATE kontrak_vendor SET sisa_pembayaran=? WHERE id_kontrak=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(temp_sisa_pembayaran, ID_Kontrak)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	path := ""

	sqlStatement = "SELECT foto_invoice FROM pembayaran_vendor WHERE id_kontrak=? "

	_ = con.QueryRow(sqlStatement, ID_Kontrak).Scan(&path)

	if path != "" {
		os.Remove(path)
	}

	sqlstatement = "DELETE FROM pembayaran_vendor WHERE id_PV=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err = stmt.Exec(obj.Id_PV)

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

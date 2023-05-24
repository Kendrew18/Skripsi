package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generate_Id_Laporan_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_LV FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_LV=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Generate_Id_Foto_Laporan_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_ft_lpv FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_ft_lpv=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//data
func Input_Laporan_Vendor(id_proyek string, id_kontrak string, laporan string, tanggal_laporan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Laporan_Vendor()

	nm_str := strconv.Itoa(nm)

	id_LV := "LV-" + nm_str

	sqlStatement := "INSERT INTO laporan_vendor (id_proyek,id_kontrak_vendor,id_laporan_vendor,laporan,tanggal_laporan,photo_laporan,status_laporan) values(?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_kontrak, id_LV, laporan, date_sql, "", 0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Laporan_Vendor(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Laporan_Vendor
	var invent str.Read_Laporan_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan_vendor, nama_vendor,pekerjaan_vendor, laporan, tanggal_laporan,status_laporan FROM laporan_vendor join kontrak_vendor on laporan_vendor.id_kontrak_vendor=kontrak_vendor.id_kontrak WHERE laporan_vendor.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan_vendor, &invent.Nama_vendor, &invent.Pekerjaan_vendor,
			&invent.Laporan, &invent.Tanggal_laporan, &invent.Status_laporan)
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

func Update_Laporan_Vendor(id_laporan_vendor string, id_kontrak string, laporan string, tanggal_laporan string) (tools.Response, error) {

	var res tools.Response
	var st str.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&st.Status)

	if st.Status == 0 {

		sqlStatement = "UPDATE laporan_vendor SET id_kontrak_vendor=?,laporan=?,tanggal_laporan=? WHERE id_laporan_vendor=?"

		date, _ := time.Parse("02-01-2006", tanggal_laporan)
		date_sql := date.Format("2006-01-02")

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_kontrak, laporan, date_sql, id_laporan_vendor)

		if err != nil {
			return res, err
		}

		rowschanged, err := result.RowsAffected()

		if err != nil {
			return res, err
		}

		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": rowschanged,
		}
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
	}

	return res, nil
}

func Update_Status_Laporan_Vendor(id_laporan_vendor string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE laporan_vendor SET status_laporan=? WHERE id_laporan_vendor=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_laporan_vendor)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//foto
func Upload_Foto_laporan_vendor(id_laporan_vendor string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response
	var foto str.Foto

	con := db.CreateCon()

	nm := Generate_Id_Foto_Laporan_Vendor()

	nm_str := strconv.Itoa(nm)

	id_LPV_FT := id_laporan_vendor + "-LPV-FT-" + nm_str

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
		path = "uploads/" + id_LPV_FT + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_LPV_FT + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_LPV_FT + ".png"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.png")
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

	sqlStatement := "SELECT photo_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_LPV_FT).Scan(&foto.Path_foto)

	foto.Path_foto = foto.Path_foto + "|" + path + "|"

	sqlstatement := "UPDATE laporan SET foto_laporan=? WHERE id_laporan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(foto.Path_foto, id_LPV_FT)

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

func Read_Foto_Laporan_Vendor(id_laporan_vendor string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Foto
	var invent str.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT tanggal_laporan FROM laporan_vendor WHERE id_laporan_vendor=? "

	err := con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&invent.Path_foto)

	if err != nil {
		return res, err
	}

	foto_sp := tools.String_Separator_To_String(invent.Path_foto)

	for i := 0; i < len(foto_sp); i++ {
		invent.Path_foto = foto_sp[i]
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

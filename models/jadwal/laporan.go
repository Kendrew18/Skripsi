package jadwal

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/jadwal"
	"Skripsi/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generate_Id_Laporan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_LP FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_LP=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Generate_Id_Foto_Laporan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_LP FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_LP=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//data
func Input_Laporan(id_proyek string, laporan string, tanggal_laporan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Laporan()

	nm_str := strconv.Itoa(nm)

	id_LP := "LP-" + nm_str

	sqlStatement := "INSERT INTO laporan (id_proyek,id_laporan,laporan,tanggal_laporan,foto_laporan,status_laporan) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_LP, laporan, date_sql, "", 0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Laporan(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []jadwal.Read_Laporan
	var invent jadwal.Read_Laporan

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan, laporan, tanggal_laporan,status_laporan FROM laporan WHERE laporan.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan, &invent.Laporan, &invent.Tanggal_laporan, &invent.Status_laporan)
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

func Update_Laporan(id_laporan string, laporan string, tanggal_laporan string) (tools.Response, error) {

	var res tools.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan WHERE id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&st.Status)

	if st.Status == 0 {

		date, _ := time.Parse("02-01-2006", tanggal_laporan)
		date_sql := date.Format("2006-01-02")

		sqlStatement = "UPDATE laporan SET laporan=?,tanggal_laporan=? WHERE id_laporan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, date_sql, id_laporan)

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

func Update_Status_Laporan(id_laporan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE laporan SET status_laporan=? WHERE id_laporan=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_laporan)

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
func Upload_Foto_laporan(id_laporan string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response
	var foto str.Foto

	con := db.CreateCon()

	nm := Generate_Id_Foto_Laporan()

	nm_str := strconv.Itoa(nm)

	id_LP_FT := id_laporan + "-LP-FT-" + nm_str

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
		path = "uploads/" + id_LP_FT + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_LP_FT + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_LP_FT + ".png"
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

	sqlStatement := "SELECT foto_laporan FROM laporan WHERE laporan.id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&foto.Path_foto)

	foto.Path_foto = foto.Path_foto + "|" + path + "|"

	sqlstatement := "UPDATE laporan SET foto_laporan=? WHERE id_laporan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(foto.Path_foto, id_laporan)

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

func Read_Foto_Laporan(id_laporan string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Foto
	var invent str.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT foto_laporan FROM laporan WHERE id_laporan=? "

	err := con.QueryRow(sqlStatement, id_laporan).Scan(&invent.Path_foto)

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

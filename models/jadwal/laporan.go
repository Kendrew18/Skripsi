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

//Gerenate_Id_Laporan
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

//Gerenate_Id_Foto_Laporan
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

//Input-Laporan (done)
func Input_Laporan(id_proyek string, laporan string, tanggal_laporan string,
	id_penjadwalan string, check string) (tools.Response, error) {
	var res tools.Response
	var RP jadwal.Progress

	con := db.CreateCon()

	nm := Generate_Id_Laporan()

	nm_str := strconv.Itoa(nm)

	id_LP := "LP-" + nm_str

	sqlStatement := "INSERT INTO laporan (id_proyek,id_laporan,laporan,tanggal_laporan,foto_laporan,status_laporan,id_penjadwalan,check_box) values(?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	ip := tools.String_Separator_To_String(id_penjadwalan)
	ck := tools.String_Separator_To_Int(check)

	//masih salah kurang pengecekan udah finish ta blm e
	for i := 0; i < len(ip); i++ {
		if ck[i] == 1 {

			sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_penjadwalan, &RP.Progress,
				&RP.Durasi, &RP.Complate)

			if RP.Durasi == RP.Progress+1 {
				RP.Complate = 1
			} else {
				RP.Progress = RP.Progress + 1
			}

			sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(RP.Progress, RP.Complate, ip[i])

		} else if ck[i] == 2 {

			sqlStatement = "UPDATE penjadwalan SET complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(1, ip[i])
		}

	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_LP, laporan, date_sql, "", 0, id_penjadwalan, check)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read-Laporan (done)
func Read_Laporan(id_Proyek string) (tools.Response, error) {
	var res tools.Response

	var arr_lp []jadwal.Read_Laporan
	var lp jadwal.Read_Laporan

	var dl jadwal.Detail_Laporan

	var invent jadwal.Read_Laporan_String

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan, laporan, tanggal_laporan,status_laporan,id_penjadwalan FROM laporan WHERE laporan.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&lp.Id_laporan, &lp.Laporan, &lp.Tanggal_laporan, &lp.Status_laporan,
			&invent.Id_Penjadwalan)

		idj := tools.String_Separator_To_String(invent.Id_Penjadwalan)

		for i := 0; i < len(idj); i++ {

			durasi := 0
			complate := 0

			sqlstatemen_jdl := "SELECT id_penjadwalan,nama_task,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlstatemen_jdl, idj[i]).Scan(&dl.Id_Penjadwalan,
				&dl.Nama_Sub_Pekerjaan, &dl.Progress, &durasi, &complate)

			if complate == 1 && durasi != dl.Progress {
				dl.Progress = durasi
			}

			dl.Progress = (dl.Progress / durasi) * 100
			lp.Detail_Laporan = append(lp.Detail_Laporan, dl)

		}

		if err != nil {
			return res, err
		}
		arr_lp = append(arr_lp, lp)
	}

	if arr_lp == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_lp
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_lp
	}

	return res, nil
}

//Update-Laporan (done)
func Update_Laporan(id_laporan string, laporan string, tanggal_laporan string, id_penjadwalan string, check string) (tools.Response, error) {

	var res tools.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan WHERE id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&st.Status)

	if st.Status == 0 {

		var read_dt_lp jadwal.Read_Laporan_String
		var rp jadwal.Progress
		var RP jadwal.Progress
		temp := ""

		sqlStatement := "SELECT id_penjadwalan,check_box FROM laporan WHERE id_laporan=?"

		_ = con.QueryRow(sqlStatement, id_laporan).Scan(&read_dt_lp.Id_Penjadwalan, &temp)

		id := tools.String_Separator_To_String(read_dt_lp.Id_Penjadwalan)

		for i := 0; i < len(id); i++ {
			sqlStatement = "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlStatement, id[i]).Scan(&rp.Id_penjadwalan,
				&rp.Progress, rp.Durasi, rp.Complate)

			if rp.Complate == 1 {
				rp.Complate = 0
			} else if rp.Complate == 0 {
				rp.Progress--
			}
		}

		id_br := tools.String_Separator_To_String(id_penjadwalan)
		ck := tools.String_Separator_To_Int(check)

		for i := 0; i < len(id_br); i++ {
			if ck[i] == 1 {

				sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

				_ = con.QueryRow(sqlstatemen_jdl, id_br[i]).Scan(&RP.Id_penjadwalan, &RP.Progress,
					&RP.Durasi, &RP.Complate)

				if RP.Durasi == RP.Progress+1 {
					RP.Complate = 1
				} else {
					RP.Progress = RP.Progress + 1
				}

				sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(RP.Progress, RP.Complate, id_br[i])

			} else if ck[i] == 2 {

				sqlStatement = "UPDATE penjadwalan SET complate=? WHERE id_penjadwalan=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(1, id_br[i])
			}
		}

		date, _ := time.Parse("02-01-2006", tanggal_laporan)
		date_sql := date.Format("2006-01-02")

		sqlStatement = "UPDATE laporan SET laporan=?,tanggal_laporan=?,id_penjadwalan=?,check_box=? WHERE id_laporan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, date_sql, id_penjadwalan, check, id_laporan)

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

//Update-Status-Laporan (done)
func Update_Status_Laporan(id_laporan string) (tools.Response, error) {
	var res tools.Response
	var RP jadwal.Progress

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

	id_jdl := ""

	sqlStatement = "SELECT id_penjadwalan FROM laporan WHERE id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&id_jdl)

	id := tools.String_Separator_To_String(id_jdl)

	for i := 0; i < len(id); i++ {
		sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

		_ = con.QueryRow(sqlstatemen_jdl, id[i]).Scan(&RP.Id_penjadwalan, &RP.Progress,
			&RP.Durasi, &RP.Complate)

		if RP.Complate == 1 {
			RP.Progress = RP.Durasi
		}

		sqlStatement = "UPDATE penjadwalan SET progress=? WHERE id_penjadwalan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(RP.Progress, id[i])
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Upload-Foto-Laporan (done)
func Upload_Foto_Laporan(id_laporan string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
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

//Read-Foto-Laporan (done)
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

//See-Task-Di-Input-Laporan (done)
func See_Task(tanggal_laporan string) (tools.Response, error) {
	var res tools.Response

	var rt_lp jadwal.Read_Task_Laporan
	var arr_rt_lp []jadwal.Read_Task_Laporan

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,nama_task,durasi,progress FROM penjadwalan WHERE tanggal_dimulai<=? && tanggal_selesai>=? && penjadwalan.progress != penjadwalan.durasi"

	rows, err := con.Query(sqlStatement, date_sql)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		durasi := 0

		err = rows.Scan(&rt_lp.Id_penjadwalan, &rt_lp.Nama_Task, &durasi, &rt_lp.Progress)

		rt_lp.Progress = (rt_lp.Progress / durasi) * 100

		if err != nil {
			return res, err
		}

		arr_rt_lp = append(arr_rt_lp, rt_lp)
	}

	if arr_rt_lp == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_rt_lp
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_rt_lp
	}

	return res, nil
}

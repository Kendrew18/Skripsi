package vendor_all

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/jadwal"
	"Skripsi/struct_all/vendor_all"
	"Skripsi/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Generate-id-laporan-vendor(done)
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

//Generate-id-foto-vendor(done)
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

//input-laporan-Vendor (done)
func Input_Laporan_Vendor(id_proyek string, laporan string, tanggal_laporan string, id_kontrak string, check_box string) (tools.Response, error) {
	var res tools.Response
	var RP vendor_all.Progress_Vendor

	con := db.CreateCon()

	nm := Generate_Id_Laporan_Vendor()

	nm_str := strconv.Itoa(nm)

	id_LV := "LV-" + nm_str

	sqlStatement := "INSERT INTO laporan_vendor (id_proyek,id_laporan_vendor,laporan,tanggal_laporan,photo_laporan,status_laporan,id_kontrak,check_box) values(?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	ip := tools.String_Separator_To_String(id_kontrak)
	ck := tools.String_Separator_To_Int(check_box)

	for i := 0; i < len(ip); i++ {
		if ck[i] == 1 {

			sqlstatemen_jdl := "SELECT id_kontrak,working_progress,working_complate,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) FROM kontrak_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess,
				&RP.Durasi, &RP.Working_Complate)

			if RP.Durasi == RP.Working_Progess+1 {
				RP.Working_Complate = 1
			} else {
				RP.Working_Progess = RP.Working_Progess + 1
			}

			sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(RP.Working_Progess, RP.Working_Complate, ip[i])

		} else if ck[i] == 2 {

			sqlStatement = "UPDATE kontrak_vendor SET working_complate=? WHERE id_kontrak=?"

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

	_, err = stmt.Exec(id_proyek, id_LV, laporan, date_sql, "", 0, id_kontrak, check_box)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//read-laporan-Vendor(done)
func Read_Laporan_Vendor(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Read_Laporan_Vendor
	var invent vendor_all.Read_Laporan_Vendor
	var RlV vendor_all.Read_Laporan_Vendor_String
	var dl vendor_all.Detail_Laporan_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan_vendor, laporan, tanggal_laporan,status_laporan,id_kontrak,check_box FROM laporan_vendor WHERE laporan_vendor.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan_vendor, &invent.Laporan,
			&invent.Tanggal_laporan, &invent.Status_laporan, &RlV.Id_Kontrak_Vendor)
		kv := tools.String_Separator_To_String(RlV.Id_Kontrak_Vendor)

		for i := 0; i < len(kv); i++ {
			durasi := 0
			complate := 0

			sqlstatemen_jdl := "SELECT id_kontrak,nama_vendor,penkerjaan_vendor,working_progress,datediff(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor JOIN vendor v on kontrak_vendor.id_MV=v.id_master_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlstatemen_jdl, kv[i]).Scan(&dl.Id_kontrak,
				&dl.Nama_vendor, &dl.Pekerjaan_vendor, &dl.Progress, &durasi, &complate)

			if complate == 1 && durasi != dl.Progress {
				dl.Progress = durasi
			}

			dl.Progress = (dl.Progress / durasi) * 100
			invent.Detail_Laporan_Vendor = append(invent.Detail_Laporan_Vendor, dl)
		}

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

//Update-laporan-Vendor (done)
func Update_Laporan_Vendor(id_laporan_vendor string, laporan string, tanggal_laporan string, id_kontrak string, check_box string) (tools.Response, error) {

	var res tools.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&st.Status)

	if st.Status == 0 {

		var read_dt_lp vendor_all.Read_Laporan_Vendor_String
		var rp vendor_all.Progress_Vendor
		var RP vendor_all.Progress_Vendor
		temp := ""

		sqlStatement := "SELECT id_kontrak,check_box FROM laporan_vendor WHERE id_laporan_vendor=?"

		_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&read_dt_lp.Id_Kontrak_Vendor, &temp)

		id := tools.String_Separator_To_String(read_dt_lp.Id_Kontrak_Vendor)

		for i := 0; i < len(id); i++ {
			sqlStatement = "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlStatement, id[i]).Scan(&rp.Id_kontrak,
				&rp.Working_Progess, rp.Durasi, rp.Working_Complate)

			if rp.Working_Complate == 1 {
				rp.Working_Complate = 0
			} else if rp.Working_Complate == 0 {
				rp.Working_Progess--
			}
		}

		ip := tools.String_Separator_To_String(id_kontrak)
		ck := tools.String_Separator_To_Int(check_box)

		for i := 0; i < len(ip); i++ {
			if ck[i] == 1 {

				sqlstatemen_jdl := "SELECT id_kontrak,working_progress,working_complate,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) FROM kontrak_vendor WHERE id_kontrak=?"

				_ = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess,
					&RP.Durasi, &RP.Working_Complate)

				if RP.Durasi == RP.Working_Progess+1 {
					RP.Working_Complate = 1
				} else {
					RP.Working_Progess = RP.Working_Progess + 1
				}

				sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(RP.Working_Progess, RP.Working_Complate, ip[i])

			} else if ck[i] == 2 {

				sqlStatement = "UPDATE kontrak_vendor SET working_complate=? WHERE id_kontrak=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(1, ip[i])
			}

		}

		sqlStatement = "UPDATE laporan_vendor SET laporan=?,tanggal_laporan=?,id_kontrak=?,check_box=? WHERE id_laporan_vendor=?"

		date, _ := time.Parse("02-01-2006", tanggal_laporan)
		date_sql := date.Format("2006-01-02")

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, date_sql, id_kontrak, check_box, id_laporan_vendor)

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

//Update-Status-laporan-Vendor(done)
func Update_Status_Laporan_Vendor(id_laporan_vendor string) (tools.Response, error) {
	var res tools.Response
	var RP vendor_all.Progress_Vendor

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

	id_jdl := ""

	sqlStatement = "SELECT id_kontrak FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&id_jdl)

	id := tools.String_Separator_To_String(id_jdl)

	for i := 0; i < len(id); i++ {
		sqlstatemen_jdl := "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

		_ = con.QueryRow(sqlstatemen_jdl, id[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess,
			&RP.Durasi, &RP.Working_Complate)

		if RP.Working_Complate == 1 {
			RP.Working_Progess = RP.Durasi
		}

		sqlStatement = "UPDATE kontrak_vendor SET working_progress=? WHERE id_kontrak=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(RP.Working_Progess, id[i])
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Upload_Foto_Laporan(done)
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

//Read_Foto_Laporan(done)
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

//See-Task-Di-Input-Laporan (done)
func See_Task_Vendor(tanggal_laporan string) (tools.Response, error) {
	var res tools.Response

	var rt_lp vendor_all.Read_Task_Laporan_Vendor
	var arr_rt_lp []vendor_all.Read_Task_Laporan_Vendor

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	con := db.CreateCon()

	sqlStatement := "SELECT id_kontrak,nama_vendor,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_progress FROM kontrak_vendor JOIN vendor v ON v.id_master_vendor=kontrak_vendor.id_MV WHERE tanggal_pengerjaan_dimulai<=? && tanggal_pengerjaan_berakhir>=? && working_progress != DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai)"

	rows, err := con.Query(sqlStatement, date_sql)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		durasi := 0

		err = rows.Scan(&rt_lp.Id_Kontrak, &rt_lp.Nama_Vendor, &durasi, &rt_lp.Pekerjaan_Vendor, &rt_lp.Progress)

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

package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
	"net/http"
	"strconv"
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

func Input_Laporan_Vendor(id_proyek string, id_kontrak string, laporan string, tanggal_laporan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Laporan_Vendor()

	nm_str := strconv.Itoa(nm)

	id_LV := "LV-" + nm_str

	sqlStatement := "INSERT INTO laporan_vendor (id_proyek,id_kontrak_vendor,id_laporan_vendor,laporan,tanggal_laporan,photo_laporan) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_kontrak, id_LV, laporan, date_sql, "|images.png|")

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

	sqlStatement := "SELECT id_laporan_vendor, nama_vendor,pekerjaan_vendor, laporan, tanggal_laporan, photo_laporan FROM laporan_vendor join kontrak_vendor on laporan_vendor.id_kontrak_vendor=kontrak_vendor.id_kontrak WHERE laporan_vendor.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan_vendor, &invent.Nama_vendor, &invent.Pekerjaan_vendor,
			&invent.Laporan, &invent.Tanggal_laporan, &invent.Foto_laporan)
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

func Update_Laporan_Vendor(id_laporan_vendor string, laporan string) (tools.Response, error) {

	var res tools.Response
	var st str.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&st.Status)

	if st.Status == 0 {

		sqlStatement = "UPDATE laporan_vendor SET laporan=? WHERE id_laporan_vendor=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, id_laporan_vendor)

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

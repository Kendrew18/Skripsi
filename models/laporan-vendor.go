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

	sqlStatement := "SELECT id_kv FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_kv=?"

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
	var arr_invent []str.Read_Kontrak_Vendor
	var invent str.Read_Kontrak_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT  id_laporan_vendor, nama_vendor,jenis_pekerjaan, laporan, tanggal_laporan, photo_laporan FROM laporan_vendor join kontrak_vendor on laporan_vendor.id_kontrak_vendor=kontrak_vendor.id_kontrak WHERE laporan_vendor.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_proyek, &invent.Id_kontak, &invent.Nomor_Kontrak, &invent.Nama_vendor,
			&invent.Total_nilai_kontrak, &invent.Nomial_Pembayaran, &invent.Jenis_Pekerjaan,
			&invent.Tanggal_mulai_kontrak, &invent.Tanggal_berakhir_kontrak, &invent.Sisa_pembayaran)
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

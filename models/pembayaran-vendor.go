package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
	"net/http"
	"strconv"
	"time"
)

func Generate_Id_Pembayaran_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_pv FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_pv=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Pembayaran_Vendor(id_kontrak string, nomor_invoice string,
	jumlah_pembayaran int64, tanggal_pembayaran string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Kotrak_Vendor()

	nm_str := strconv.Itoa(nm)

	id_PV := "PV-" + nm_str

	foto := "uploads/images.png"

	sqlStatement := "INSERT INTO pembayaran_vendor (id_PV,id_kontrak,nomor_invoice,jumlah_pembayaran,tanggal_pembayaran ,foto_invoice) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_PV, id_kontrak, nomor_invoice, jumlah_pembayaran, date_sql, foto)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Pembayaran_Vendor(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Pembayaran_Vendor
	var invent str.Read_Pembayaran_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_PV, nomor_invoice, jumlah_pembayaran, tanggal_pembayaran,foto_invoice FROM pembayaran_vendor WHERE id_Proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

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

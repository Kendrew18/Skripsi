package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Generate_Id_Kotrak_Vendor() int {
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

func Input_Kontrak_Vendor(id_proyek string, nomor_kontrak string, nama_vendor string,
	total_nilai_kontrak int64, jenis_pekerjaan_vendor string, pekerjaan_vendor string, tanggal_dimulai string,
	tanggal_selesai string, date_pengiriman string, date_dimulai string, date_selesai string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Kotrak_Vendor()

	nm_str := strconv.Itoa(nm)

	id_kontrak := "KV-" + nm_str

	sqlStatement := "INSERT INTO kontrak_vendor (id_proyek,id_kontrak,nomor_kontrak,nama_vendor,total_nilai_kontrak ,nominal_pembayaran,jenis_pekerjaan_vendor,pekerjaan_vendor,tanggal_mulai_kontrak,tanggal_berakhir_kontrak,sisa_pembayaran,tanggal_pengiriman,tanggal_pengerjaan_dimulai,tanggal_pengerjaan_berakhir) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_dimulai)
	date_sql := date.Format("2006-01-02")
	dm := date.Format("200601")

	date2, _ := time.Parse("02-01-2006", tanggal_selesai)
	date_sql2 := date2.Format("2006-01-02")
	dt := date2.Format("200601")

	date3, _ := time.Parse("02-01-2006", date_pengiriman)
	date_sql3 := date3.Format("2006-01-02")

	date4, _ := time.Parse("02-01-2006", date_dimulai)
	date_sql4 := date4.Format("2006-01-02")

	date5, _ := time.Parse("02-01-2006", date_selesai)
	date_sql5 := date5.Format("2006-01-02")

	fmt.Println(dm)
	fmt.Println(dt)

	sqlStatement2 := "SELECT period_diff( " + dt + ", " + dm + ")"

	var temp int64

	temp = 0

	_ = con.QueryRow(sqlStatement2).Scan(&temp)

	var nominal_pembayaran int64

	nominal_pembayaran = 0

	if temp == 0 {
		nominal_pembayaran = total_nilai_kontrak
	} else {
		nominal_pembayaran = total_nilai_kontrak / temp
	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_kontrak, nomor_kontrak, nama_vendor, total_nilai_kontrak,
		nominal_pembayaran, jenis_pekerjaan_vendor, pekerjaan_vendor, date_sql, date_sql2,
		total_nilai_kontrak, date_sql3, date_sql4, date_sql5)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Kontrak_Vendor(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Kontrak_Vendor
	var invent str.Read_Kontrak_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek,id_kontrak,nomor_kontrak,nama_vendor,total_nilai_kontrak,nominal_pembayaran,pekerjaan_vendor,tanggal_mulai_kontrak,tanggal_berakhir_kontrak,sisa_pembayaran FROM kontrak_vendor WHERE id_Proyek=? ORDER BY co ASC "

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

func Delete_Kontrak_Vendor(id_kontrak string) (tools.Response, error) {
	var res tools.Response
	var arrobj []str.Read_id_pv
	var obj str.Read_id_pv

	con := db.CreateCon()

	sqlStatement := "SELECT id_PV FROM pembayaran_vendor WHERE id_kontrak=? "

	rows, err := con.Query(sqlStatement, id_kontrak)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_pv)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	if arrobj == nil {

		sqlstatement := "DELETE FROM kontrak_vendor WHERE id_kontrak=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_kontrak)

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

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Tidak bisa di hapus"
	}

	return res, nil
}

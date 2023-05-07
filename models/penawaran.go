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

func Generate_Id_Header_Penawaran() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_hp FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_hp=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Header_Penawaran(id_proyek string, kode_surat string, tanggal_dibuat string,
	nama_perusahaan string, alamat_perusahaan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Header_Penawaran()

	nm_str := strconv.Itoa(nm)

	id_kontrak := "HP-" + nm_str

	sqlStatement := "INSERT INTO header_penawaran (id_proyek,id_header_penawaran,kode_surat,tanggal_dibuat,nama_perusahaan,alamat_perusahaan) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_dibuat)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_kontrak, kode_surat, date_sql, nama_perusahaan, alamat_perusahaan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Header_Penawaran(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Header
	var invent str.Read_Header

	con := db.CreateCon()

	sqlStatement := "SELECT  id_header_penawaran, kode_surat, tanggal_dibuat, nama_perusahaan, alamat_perusahaan FROM header_penawaran WHERE id_Proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_Header_Penawwaran, &invent.Kode_surat, &invent.Tanggal_dibuat, &invent.Nama_Perusahaan,
			&invent.Alamat_Perusahaan)
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

func Input_Penawaran(id_proyek string, nomor_kontrak string, nama_vendor string,
	total_nilai_kontrak int64, jenis_pekerjaan string, tanggal_dimulai string,
	tanggal_selesai string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Kotrak_Vendor()

	nm_str := strconv.Itoa(nm)

	id_kontrak := "KV-" + nm_str

	sqlStatement := "INSERT INTO kontrak_vendor (id_proyek,id_kontrak,nomor_kontrak,nama_vendor,total_nilai_kontrak ,nominal_pembayaran,jenis_pekerjaan,tanggal_mulai_kontrak,tanggal_berakhir_kontrak,sisa_pembayaran) values(?,?,?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_dimulai)
	date_sql := date.Format("2006-01-02")
	dm := date.Format("200601")

	date2, _ := time.Parse("02-01-2006", tanggal_selesai)
	date_sql2 := date2.Format("2006-01-02")
	dt := date2.Format("200601")

	fmt.Println(dm)
	fmt.Println(dt)

	sqlStatement2 := "SELECT period_diff( " + dt + ", " + dm + ")"

	var temp int64

	temp = 0

	_ = con.QueryRow(sqlStatement2).Scan(&temp)

	nominal_pembayaran := total_nilai_kontrak / temp

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_kontrak, nomor_kontrak, nama_vendor, total_nilai_kontrak,
		nominal_pembayaran, jenis_pekerjaan, date_sql, date_sql2, total_nilai_kontrak)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Penawaran(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Kontrak_Vendor
	var invent str.Read_Kontrak_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek,id_kontrak,nomor_kontrak,nama_vendor,total_nilai_kontrak,nominal_pembayaran,jenis_pekerjaan,tanggal_mulai_kontrak,tanggal_berakhir_kontrak,sisa_pembayaran FROM kontrak_vendor WHERE id_Proyek=? ORDER BY co ASC "

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

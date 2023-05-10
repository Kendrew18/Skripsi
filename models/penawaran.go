package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
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

func Generate_Id_Penawaran() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT penawaran FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET penawaran=?"

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

func Input_Penawaran(id_proyek string, judul string, sub_pekerjaan string,
	keterangan string, jumlah string, satuan string,
	harga string, total string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Penawaran()

	nm_str := strconv.Itoa(nm)

	id_penawaran := "P-" + nm_str

	sqlStatement := "INSERT INTO penawaran (id_penawaran,id_proyek,judul,sub_pekerjaan,keterangan,jumlah,satuan,harga,total,sub_total) values(?,?,?,?,?,?,?,?,?,?)"

	ttl := tools.String_Separator_To_Int64(total)

	var sub_total int64
	sub_total = 0

	for i := 0; i < len(ttl); i++ {
		sub_total += ttl[i]
	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_penawaran, id_proyek, judul, sub_pekerjaan, keterangan, jumlah,
		satuan, harga, total, sub_total)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Penawaran(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Penawaran
	var invent str.Read_Penawaran
	var tmp str.Read_Penawaran_String

	con := db.CreateCon()

	sqlStatement := "SELECT judul, sub_pekerjaan, keterangan, jumlah, satuan, harga, total, sub_total FROM penawaran WHERE id_Proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Judul, &tmp.Sub_pekerjaan, &tmp.Keterangan, &tmp.Jumlah, &tmp.Satuan,
			&tmp.Harga, &tmp.Total, &invent.Sub_total)

		invent.Sub_pekerjaan = tools.String_Separator_To_String(tmp.Sub_pekerjaan)
		invent.Keterangan = tools.String_Separator_To_String(tmp.Keterangan)
		invent.Jumlah = tools.String_Separator_To_float64(tmp.Jumlah)
		invent.Satuan = tools.String_Separator_To_String(tmp.Satuan)
		invent.Harga = tools.String_Separator_To_Int64(tmp.Harga)
		invent.Total = tools.String_Separator_To_Int64(tmp.Total)

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

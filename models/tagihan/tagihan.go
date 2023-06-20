package tagihan

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/tagihan"
	"Skripsi/tools"
	"net/http"
	"strconv"
	"time"
)

//Generate Id Tagihan
func Generate_Id_Tagihan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_tagihan FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_tagihan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input-Tagihan
func Input_Tagihan(id_proyek string, perihal string, tanggal_pembarian_kwitansi string,
	tanggal_pembayaran string, nominal_keseluruhan int64, id_penawaran string,
	id_sub_pekerjaan string, nominal string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Tagihan()

	nm_str := strconv.Itoa(nm)

	id_real := "TG-" + nm_str

	sqlStatement := "INSERT INTO tagihan (id_tagihan,id_proyek, perihal_tagihan, tanggal_pemberian_kwitansi, tanggal_pembayaran, nominal_keseluruhan,id_penawaran, id_sub_pekerjaan, nominal) values(?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-2006", tanggal_pembarian_kwitansi)
	date_sql2 := date2.Format("2006-01-02")

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_real, id_proyek, perihal, date_sql2, date_sql,
		nominal_keseluruhan, id_penawaran, id_sub_pekerjaan, nominal)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read-Tagihan
func Read_Realisasi(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []tagihan.Read_Tagihan
	var invent tagihan.Read_Tagihan

	var rtf tagihan.Read_Detail_Tagihan

	con := db.CreateCon()

	sqlStatement := "SELECT id_tagihan,perihal_tagihan, DATE_FORMAT(tanggal_pemberian_kwitansi, '%d-%m%-%Y'),DATE_FORMAT(tanggal_pembayaran, '%d-%m%-%Y'),nominal_keseluruhan,id_penawaran,id_sub_pekerjaan,nominal FROM tagihan WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		ip := ""
		isp := ""
		nm := ""
		err = rows.Scan(&invent.Id_Tagihan, &invent.Perihal_Tagihan, &invent.Tanggal_Pemberian_Kwitansi,
			&invent.Tanggal_Pembayaran, &invent.Nominal_Keseluruhan, &ip, &isp, &nm)

		temp_ip := tools.String_Separator_To_String(ip)
		temp_isp := tools.String_Separator_To_String(isp)
		temp_nm := tools.String_Separator_To_Int64(nm)

		for i := 0; i < len(temp_isp); i++ {
			rtf.Id_Penawaran = temp_ip[i]
			rtf.Id_Sub_Pekerjaan = temp_isp[i]
			rtf.Nominal = temp_nm[i]

			sqlStatement = "SELECT nama_task FROM  penjadwalan WHERE id_penawaran=? && id_sub_pekerjaan=?"

			_ = con.QueryRow(sqlStatement, rtf.Id_Penawaran, rtf.Id_Sub_Pekerjaan).Scan(
				&rtf.Nama_Sub_Pekerjaan)
			invent.Read_Detail_Tagihan = append(invent.Read_Detail_Tagihan, rtf)
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

//Delete-Tagihan
func Delete_Tagihan(id_tagihan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "DELETE FROM tagihan WHERE id_tagihan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_tagihan)

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

	return res, nil
}

//See-Judul
func See_Judul(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []tagihan.See_Judul
	var invent tagihan.See_Judul

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&invent.Id_Penawaran, &invent.Judul)

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

//See-Sub-Pekerjaan
func See_Sub_Pekerjaan(id_proyek string, id_penawaran string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []tagihan.See_Sub_Pekerjaan
	var invent tagihan.See_Sub_Pekerjaan

	con := db.CreateCon()

	sqlStatement := "SELECT id_sub_pekerjaan,sub_pekerjaan FROM penawaran WHERE id_proyek=? && id_penawaran=? ORDER BY co ASC "

	temp := ""
	temp2 := ""

	err := con.QueryRow(sqlStatement, id_proyek, id_penawaran).Scan(&temp, &temp2)

	if err != nil {
		return res, err
	}

	tp := tools.String_Separator_To_String(temp)
	tp2 := tools.String_Separator_To_String(temp2)

	for i := 0; i < len(tp); i++ {
		invent.Sub_Pekerjaan = tp2[i]
		invent.Id_Sub_Pekerjaan = tp[i]
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

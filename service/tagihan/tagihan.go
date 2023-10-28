package tagihan

import (
	"Skripsi/config/db"
	"Skripsi/models/tagihan"
	tools2 "Skripsi/service/tools"
	"net/http"
	"strconv"
	"time"
)

//Input-Tagihan
func Input_Tagihan(id_proyek string, perihal string, tanggal_pembarian_kwitansi string, tanggal_pembayaran string, nominal_keseluruhan int64, id_penawaran string, id_sub_pekerjaan string, nominal string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM tagihan ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_Tagihan := "TG-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO tagihan (co,id_tagihan,id_proyek, perihal_tagihan, tanggal_pemberian_kwitansi, tanggal_pembayaran, nominal_keseluruhan) values(?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-2006", tanggal_pembarian_kwitansi)
	date_sql2 := date2.Format("2006-01-02")

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_Tagihan, id_proyek, perihal, date_sql2, date_sql, nominal_keseluruhan)

	if err != nil {
		return res, err
	}

	id_p := tools2.String_Separator_To_String(id_penawaran)
	id_s_p := tools2.String_Separator_To_String(id_sub_pekerjaan)
	nom := tools2.String_Separator_To_Int64(nominal)

	for i := 0; i < len(id_p); i++ {
		nm_str2 := 0

		Sqlstatement = "SELECT co FROM detail_tagihan ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str2)

		nm_str2 = nm_str2 + 1

		id_Detail_Tagihan := "DTG-" + strconv.Itoa(nm_str2)

		sqlStatement = "INSERT INTO detail_tagihan (co, id_detail_tagihan, id_tagihan, id_penawaran, id_sub_pekerjaan, nominal) values(?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str2, id_Detail_Tagihan, id_Tagihan, id_p[i], id_s_p[i], nom[i])

		if err != nil {
			return res, err
		}
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read-Tagihan
func Read_Tagihan(id_proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []tagihan.Read_Tagihan
	var invent tagihan.Read_Tagihan

	con := db.CreateCon()

	sqlStatement := "SELECT id_tagihan,perihal_tagihan, DATE_FORMAT(tanggal_pemberian_kwitansi, '%d-%m%-%Y'),DATE_FORMAT(tanggal_pembayaran, '%d-%m%-%Y'),nominal_keseluruhan FROM tagihan WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var obj tagihan.Read_Detail_Tagihan
		var arr_obj []tagihan.Read_Detail_Tagihan
		err = rows.Scan(&invent.Id_Tagihan, &invent.Perihal_Tagihan, &invent.Tanggal_Pemberian_Kwitansi, &invent.Tanggal_Pembayaran, &invent.Nominal_Keseluruhan)

		sqlStatement := "SELECT id_detail_tagihan, detail_tagihan.id_penawaran, detail_tagihan.id_sub_pekerjaan,nama_sub_pekerjaan, nominal FROM detail_tagihan JOIN detail_penawaran dp on detail_tagihan.id_sub_pekerjaan = dp.id_sub_pekerjaan WHERE id_tagihan=? ORDER BY detail_tagihan.co ASC "

		rows2, err := con.Query(sqlStatement, invent.Id_Tagihan)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			err = rows2.Scan(&obj.Id_detail_tagihan, &obj.Id_Penawaran, &obj.Id_Sub_Pekerjaan, &obj.Nama_Sub_Pekerjaan, &obj.Nominal)

			if err != nil {
				return res, err
			}
			arr_obj = append(arr_obj, obj)
		}

		invent.Read_Detail_Tagihan = arr_obj

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
func Delete_Tagihan(id_tagihan string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	id := ""
	sqlstate := "SELECT id_detail_tagihan FROM detail_tagihan WHERE id_tagihan=?"

	err := con.QueryRow(sqlstate, id_tagihan).Scan(&id)

	if id != "" {
		sqlstatement := "DELETE FROM detail_tagihan WHERE id_tagihan=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_tagihan)

		if err != nil {
			return res, err
		}

		_, err = result.RowsAffected()

		if err != nil {
			return res, err
		}
	}

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
func See_Judul(id_proyek string) (tools2.Response, error) {
	var res tools2.Response
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
func See_Sub_Pekerjaan(id_penawaran string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []tagihan.See_Sub_Pekerjaan
	var invent tagihan.See_Sub_Pekerjaan

	con := db.CreateCon()

	sqlStatement := "SELECT id_sub_pekerjaan,nama_sub_pekerjaan FROM detail_penawaran WHERE id_penawaran=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&invent.Id_Sub_Pekerjaan, &invent.Sub_Pekerjaan)

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

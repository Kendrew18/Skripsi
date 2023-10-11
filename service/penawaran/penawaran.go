package penawaran

import (
	"Skripsi/config/db"
	"Skripsi/models/penawaran"
	tools2 "Skripsi/service/tools"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

//Input Penawaran
func Input_Penawaran(id_proyek string, judul string, sub_pekerjaan string, catatan string, jumlah string, satuan string, harga string, sub_total string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	//input penawaran
	nm_str := 0

	Sqlstatement := "SELECT co FROM penawaran ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_penawaran := "P-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO penawaran (co,id_penawaran,id_proyek,judul,total,status_penawaran) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_penawaran, id_proyek, judul, 0, 0)

	if err != nil {
		return res, err
	}

	//input sub pekerjaan dan input penjadwalan
	var total int64
	total = 0

	Sub_pekerjaan := tools2.String_Separator_To_String(sub_pekerjaan)
	Catatan := tools2.String_Separator_To_String(catatan)
	Jumlah := tools2.String_Separator_To_float64(jumlah)
	Satuan := tools2.String_Separator_To_String(satuan)
	Harga := tools2.String_Separator_To_Int64(harga)
	ttl := tools2.String_Separator_To_float64(sub_total)

	for i := 0; i < len(ttl); i++ {
		sbt_i := int64(math.Round(ttl[i]*100) / 100)
		total += sbt_i
	}

	for i := 0; i < len(ttl); i++ {

		//input penjadwalan
		nm_str := 0

		Sqlstatement := "SELECT co FROM penjadwalan ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

		nm_str = nm_str + 1

		id_penjadwalan := "PJD-" + strconv.Itoa(nm_str)

		sqlStatement = "INSERT INTO penjadwalan (co,id_penjadwalan,id_proyek,id_penawaran,nama_task,status_urutan,dependencies) values(?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str, id_penjadwalan, id_proyek, id_penawaran, Sub_pekerjaan[i], 0, "")

		//input sub pekerjaan
		nm_str_DP := 0

		Sqlstatement = "SELECT co FROM detail_penawaran ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DP)

		nm_str_DP = nm_str_DP + 1

		id_detail_penawaran := "SP-" + strconv.Itoa(nm_str_DP)

		sqlStatement = "INSERT INTO detail_penawaran (co, id_sub_pekerjaan, nama_sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sub_total, catatan) values(?,?,?,?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		sbt_i := int64(math.Round(ttl[i]*100) / 100)
		total += sbt_i

		_, err = stmt.Exec(nm_str_DP, id_detail_penawaran, Sub_pekerjaan[i], id_penawaran, Jumlah[i], Harga[i], Satuan[i], sbt_i, Catatan[i])

		if err != nil {
			return res, err
		}
	}

	sqlStatement = "UPDATE penawaran SET total=? WHERE id_penawaran=?"

	stmt, err = con.Prepare(sqlStatement)

	_, err = stmt.Exec(total, id_penawaran)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//input_sub_pekerjaan
func Input_Sub_Pekerjaan(id_proyek string, id_penawaran string, sub_pekerjaan string, catatan string, jumlah float64, satuan string, harga int64, sub_total float64) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	//insert detail penawaran
	nm_str_DP := 0

	Sqlstatement := "SELECT co FROM detail_penawaran ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DP)

	nm_str_DP = nm_str_DP + 1

	id_detail_penawaran := "SP-" + strconv.Itoa(nm_str_DP)

	sqlStatement := "INSERT INTO detail_penawaran (co, id_sub_pekerjaan, nama_sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sub_total, catatan) values(?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	sbt_i := int64(math.Round(sub_total*100) / 100)

	_, err = stmt.Exec(nm_str_DP, id_detail_penawaran, sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sbt_i, catatan)

	//insert penjadwalan
	nm_str := 0

	Sqlstatement = "SELECT co FROM penjadwalan ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_penjadwalan := "PJD-" + strconv.Itoa(nm_str)

	sqlStatement = "INSERT INTO penjadwalan (co,id_penjadwalan,id_proyek,id_penawaran,nama_task,status_urutan,dependencies) values(?,?,?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_penjadwalan, id_proyek, id_penawaran, sub_pekerjaan, 0, "")

	//update total
	var temp penawaran.Read_Detail_Penawaran

	sqlStatement = "SELECT sub_total FROM detail_penawaran WHERE id_penawaran=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	total := int64(0)

	for rows.Next() {
		err = rows.Scan(&temp.Sub_total)

		if err != nil {
			return res, err
		}
		total += temp.Sub_total
	}

	sqlstatement := "UPDATE penawaran SET total=? WHERE id_penawaran=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(total, id_penawaran)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read Penawaran
func Read_Penawaran(id_Proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []penawaran.Read_Penawaran
	var invent penawaran.Read_Penawaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul, total,status FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penawaran, &invent.Judul, &invent.Total, &invent.Status)

		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	for i := 0; i < len(arr_invent); i++ {
		var temp penawaran.Read_Detail_Penawaran
		var temp_arr []penawaran.Read_Detail_Penawaran

		sqlStatement := "SELECT id_sub_pekerjaan, nama_sub_pekerjaan, jumlah, harga, satuan, sub_total, catatan, status FROM detail_penawaran WHERE id_penawaran=? ORDER BY co ASC "

		rows, err := con.Query(sqlStatement, arr_invent[i].Id_penawaran)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&temp.Id_sub_pekerjaan, &temp.Sub_pekerjaan, &temp.Jumlah, &temp.Harga, &temp.Satuan, &temp.Sub_total, &temp.Catatan, &temp.Status)

			if err != nil {
				return res, err
			}
			temp_arr = append(temp_arr, temp)
		}

		arr_invent[i].Read_Detail_Penawaran = temp_arr

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

//Update status Penawaran
func Update_Status_Penawaran(id_proyek string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE penawaran SET status_penawaran=?,status=? WHERE id_proyek=?"

	stmt, err := con.Prepare(sqlstatement)

	result, err := stmt.Exec(1, 1, id_proyek)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	sqlStatement := "SELECT id_penawaran FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}
	id_penawaran := ""
	for rows.Next() {
		err = rows.Scan(&id_penawaran)

		if err != nil {
			return res, err
		}

		sqlstatement := "UPDATE detail_penawaran SET status=? WHERE id_penawaran=?"

		stmt, err := con.Prepare(sqlstatement)

		result, err := stmt.Exec(1, id_penawaran)

		if err != nil {
			return res, err
		}

		_, err = result.RowsAffected()

		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Update judul Penawaran
func Update_Judul_Penawaran(id_penawaran string, judul string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE penawaran SET judul=? WHERE id_penawaran=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judul, id_penawaran)

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

	return res, nil
}

//update item penawaran
func Update_Item_Penawaran(id_penawaran string, id_sub_pekerjaan string, sub_pekerjaan string, catatan string, jumlah float64, satuan string, harga int64, sub_total float64) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE detail_penawaran SET nama_sub_pekerjaan=?, catatan=?, jumlah=?, satuan=?, harga=?, sub_total=? WHERE id_sub_pekerjaan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	sbt_i := int64(math.Round(sub_total*100) / 100)

	result, err := stmt.Exec(sub_pekerjaan, catatan, jumlah, satuan, harga, sbt_i, id_sub_pekerjaan)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	var temp penawaran.Read_Detail_Penawaran

	sqlStatement := "SELECT sub_total FROM detail_penawaran WHERE id_penawaran=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	total := int64(0)

	for rows.Next() {
		err = rows.Scan(&temp.Sub_total)

		if err != nil {
			return res, err
		}
		total += temp.Sub_total
	}

	sqlstatement = "UPDATE penawaran SET total=? WHERE id_penawaran=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err = stmt.Exec(total, id_penawaran)

	if err != nil {
		return res, err
	}

	rowschanged, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Input tambahan sub pekerjaan
func Input_Tambahan_Sub_Pekerjaan(id_proyek string, id_penawaran string, sub_pekerjaan string, catatan string, jumlah float64, satuan string, harga int64, sub_total float64, tanggal_pekerjaan_mulai string, durasi int) (tools2.Response, error) {
	var res tools2.Response

	//insert detail pekerjaan
	con := db.CreateCon()

	nm_str_DP := 0

	Sqlstatement := "SELECT co FROM detail_penawaran ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DP)

	nm_str_DP = nm_str_DP + 1

	id_detail_penawaran := "SP-" + strconv.Itoa(nm_str_DP)

	sqlStatement := "INSERT INTO detail_penawaran (co, id_sub_pekerjaan, nama_sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sub_total, catatan) values(?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	sbt_i := int64(math.Round(sub_total*100) / 100)

	_, err = stmt.Exec(nm_str_DP, id_detail_penawaran, sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sbt_i, catatan)

	if err != nil {
		return res, err
	}

	//update total
	var temp penawaran.Read_Detail_Penawaran

	sqlStatement = "SELECT sub_total FROM detail_penawaran WHERE id_penawaran=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	total := int64(0)

	for rows.Next() {
		err = rows.Scan(&temp.Sub_total)

		if err != nil {
			return res, err
		}
		total += temp.Sub_total
	}

	sqlstatement := "UPDATE penawaran SET total=? WHERE id_penawaran=?"

	stmt, err = con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(total, id_penawaran)

	//insert penjadwalan
	nm_str := 0

	Sqlstatement = "SELECT co FROM penjadwalan ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_penjadwalan := "PJD-" + strconv.Itoa(nm_str)

	date, _ := time.Parse("02-01-2006", tanggal_pekerjaan_mulai)
	date_sql := date.Format("2006-01-02")

	date_a, _ := time.Parse("02-01-2006", tanggal_pekerjaan_mulai)
	date_awal := date_a.AddDate(0, 0, durasi-1)

	fmt.Println(date_a)

	tanggal_Pekerjaan_Selesai := date_awal.Format("2006-01-02")

	sqlStatement = "INSERT INTO penjadwalan (co,id_penjadwalan,id_proyek,id_penawaran,nama_task,durasi,tanggal_dimulai,tanggal_selesai,status_urutan,progress) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_penjadwalan, id_proyek, id_penawaran, sub_pekerjaan, durasi, date_sql, tanggal_Pekerjaan_Selesai, -1, 0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//input tambahan pekerjaan tambah
func Input_Tambahan_Pekerjaan_Tambah(id_proyek string, judul string, sub_pekerjaan string, catatan string, jumlah string, satuan string, harga string, sub_total string, tanggal_mulai string, durasi string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	//input penawaran
	nm_str := 0

	Sqlstatement := "SELECT co FROM penawaran ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_penawaran := "P-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO penawaran (co,id_penawaran,id_proyek,judul,total,status_penawaran) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_penawaran, id_proyek, judul, 0, 0)

	//input sub pekerjaan dan input penjadwalan
	var total int64
	total = 0

	fmt.Println(tanggal_mulai)
	fmt.Println(durasi)

	Sub_pekerjaan := tools2.String_Separator_To_String(sub_pekerjaan)
	Catatan := tools2.String_Separator_To_String(catatan)
	Jumlah := tools2.String_Separator_To_float64(jumlah)
	Satuan := tools2.String_Separator_To_String(satuan)
	Harga := tools2.String_Separator_To_Int64(harga)
	ttl := tools2.String_Separator_To_float64(sub_total)
	tm := tools2.String_Separator_To_String(tanggal_mulai)
	dur := tools2.String_Separator_To_Int(durasi)

	for i := 0; i < len(ttl); i++ {

		//input penjadwalan
		nm_str := 0

		Sqlstatement := "SELECT co FROM penjadwalan ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

		nm_str = nm_str + 1

		id_penjadwalan := "PJD-" + strconv.Itoa(nm_str)

		fmt.Println(tm[i])

		date, _ := time.Parse("02-01-2006", tm[i])
		date_sql := date.Format("2006-01-02")

		date_a, _ := time.Parse("02-01-2006", tm[i])
		date_awal := date_a.AddDate(0, 0, dur[i]-1)

		tanggal_Pekerjaan_Selesai := date_awal.Format("2006-01-02")

		sqlStatement = "INSERT INTO penjadwalan (co,id_penjadwalan,id_proyek,id_penawaran,nama_task,tanggal_dimulai,tanggal_selesai,durasi,status_urutan,progress) values(?,?,?,?,?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str, id_penjadwalan, id_proyek, id_penawaran, Sub_pekerjaan[i], date_sql, tanggal_Pekerjaan_Selesai, dur[i], -1, 0)

		if err != nil {
			return res, err
		}

		//input sub pekerjaan
		nm_str_DP := 0

		Sqlstatement = "SELECT co FROM detail_penawaran ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DP)

		nm_str_DP = nm_str_DP + 1

		id_detail_penawaran := "SP-" + strconv.Itoa(nm_str_DP)

		sqlStatement = "INSERT INTO detail_penawaran (co, id_sub_pekerjaan, nama_sub_pekerjaan, id_penawaran, jumlah, harga, satuan, sub_total, catatan) values(?,?,?,?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		sbt_i := int64(math.Round(ttl[i]*100) / 100)
		total += sbt_i

		_, err = stmt.Exec(nm_str_DP, id_detail_penawaran, Sub_pekerjaan[i], id_penawaran, Jumlah[i], Harga[i], Satuan[i], sbt_i, Catatan[i])

	}

	sqlStatement = "UPDATE penawaran SET total=? WHERE id_penawaran=?"

	stmt, err = con.Prepare(sqlStatement)

	_, err = stmt.Exec(total, id_penawaran)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//pilih judul pekerjaan
func Pilih_Judul_Pekerjaan(id_proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []penawaran.Pilih_Judul_Penawaran
	var invent penawaran.Pilih_Judul_Penawaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penawaran, &invent.Judul_penawaran)
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

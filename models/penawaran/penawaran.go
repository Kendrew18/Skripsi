package penawaran

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/penawaran"
	"Skripsi/tools"
	"math"
	"net/http"
	"strconv"
	"time"
)

//Generate Id Penawaran
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

//Generate Id Sub Pekerjaan
func Generate_Id_Sub_Pekerjaan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_sub_pekerjaan FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_sub_pekerjaan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Generate Id Penjadwalan
func Generate_Id_Jadwal() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_jadwal FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_jadwal=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input Penawaran
func Input_Penawaran(id_proyek string, judul string, sub_pekerjaan string,
	catatan string, jumlah string, satuan string,
	harga string, sub_total string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Penawaran()

	nm_str := strconv.Itoa(nm)

	id_penawaran := "P-" + nm_str

	temp := tools.String_Separator_To_String(sub_pekerjaan)

	id_sub_fix := ""

	var id_pj_all []string

	for i := 0; i < len(temp); i++ {
		nm_s := Generate_Id_Sub_Pekerjaan()
		nm_p := Generate_Id_Jadwal()

		nm_str_s := strconv.Itoa(nm_s)
		nm_str_p := strconv.Itoa(nm_p)

		id_sub := "|" + "S-" + nm_str_s + "|"
		id_jdl := "PJ-" + nm_str_p
		id_pj_all = append(id_pj_all, id_jdl)
		id_sub_fix = id_sub_fix + id_sub
	}

	sqlStatement := "INSERT INTO penawaran (id_penawaran,id_proyek,judul,id_sub_pekerjaan,sub_pekerjaan,catatan,jumlah,satuan,harga,total,sub_total,status_penawaran) values(?,?,?,?,?,?,?,?,?,?,?,?)"

	ttl := tools.String_Separator_To_float64(sub_total)
	sbt := ""

	var total int64
	total = 0

	for i := 0; i < len(ttl); i++ {
		sbt_i := int64(math.Round(ttl[i]*100) / 100)
		sbt = sbt + "|" + strconv.FormatInt(sbt_i, 10) + "|"
		total += sbt_i
	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_penawaran, id_proyek, judul, id_sub_fix, sub_pekerjaan, catatan, jumlah,
		satuan, harga, sbt, total, 0)

	idsf := tools.String_Separator_To_String(id_sub_fix)

	for i := 0; i < len(temp); i++ {

		sqlStatement = "INSERT INTO penjadwalan (id_sub_pekerjaan,id_penjadwalan,id_proyek,id_penawaran,nama_task) values(?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(idsf[i], id_pj_all[i], id_proyek, id_penawaran, temp[i])

	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//input_sub_pekerjaan
func Input_Sub_Pekerjaan(id_proyek string, id_penawaran string, sub_pekerjaan string,
	catatan string, jumlah string, satuan string, harga string,
	sub_total float64) (tools.Response, error) {
	var res tools.Response

	var rd penawaran.Read_Penawaran_Input_Sub_Pekerjaan

	con := db.CreateCon()

	nm_s := Generate_Id_Sub_Pekerjaan()

	nm_str_s := strconv.Itoa(nm_s)

	id_sub := "|" + "S-" + nm_str_s + "|"

	Sqlstatement := "SELECT id_sub_pekerjaan, sub_pekerjaan,catatan,jumlah,satuan,harga,sub_total,total FROM penawaran WHERE id_penawaran=?"

	_ = con.QueryRow(Sqlstatement, id_penawaran).Scan(&rd.Id_sub_pekerjaan, &rd.Sub_pekerjaan,
		&rd.Catatan, &rd.Jumlah, &rd.Satuan, &rd.Harga, &rd.Sub_total, &rd.Total)

	sqlStatement := "UPDATE penawaran SET id_sub_pekerjaan=?,sub_pekerjaan=?,catatan=?,jumlah=?,satuan=?,harga=?,sub_total=?,total=? WHERE id_penawaran=?"

	rd.Id_sub_pekerjaan = rd.Id_sub_pekerjaan + id_sub
	rd.Sub_pekerjaan = rd.Sub_pekerjaan + "|" + sub_pekerjaan + "|"
	rd.Catatan = rd.Catatan + "|" + catatan + "|"
	rd.Jumlah = rd.Jumlah + "|" + jumlah + "|"
	rd.Satuan = rd.Satuan + "|" + satuan + "|"
	rd.Harga = rd.Harga + "|" + harga + "|"

	sbt_i := int64(math.Round(sub_total*100) / 100)
	sbt := "|" + strconv.FormatInt(sbt_i, 10) + "|"
	rd.Sub_total = rd.Sub_total + "|" + sbt + "|"

	rd.Total = rd.Total + sbt_i

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(rd.Id_sub_pekerjaan, rd.Sub_pekerjaan,
		rd.Catatan, rd.Jumlah, rd.Satuan, rd.Harga, rd.Sub_total,
		rd.Total, id_penawaran)

	nm_p := Generate_Id_Jadwal()

	nm_str_p := strconv.Itoa(nm_p)

	id_jdl := "PJ-" + nm_str_p

	sqlStatement = "INSERT INTO penjadwalan (id_sub_pekerjaan,id_penjadwalan,id_proyek,id_penawaran,nama_task) values(?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_sub, id_jdl, id_proyek, id_penawaran, sub_pekerjaan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read Penawaran
func Read_Penawaran(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []penawaran.Read_Penawaran
	var invent penawaran.Read_Penawaran
	var tmp penawaran.Read_Penawaran_String

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul, id_sub_pekerjaan, sub_pekerjaan, catatan, jumlah, satuan, harga, sub_total, total FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penawaran, &invent.Judul, &tmp.Id_sub_pekerjaan, &tmp.Sub_pekerjaan, &tmp.Keterangan, &tmp.Jumlah, &tmp.Satuan,
			&tmp.Harga, &tmp.Sub_total, &invent.Total)

		invent.Id_sub_pekerjaan = tools.String_Separator_To_String(tmp.Id_sub_pekerjaan)
		invent.Sub_pekerjaan = tools.String_Separator_To_String(tmp.Sub_pekerjaan)
		invent.Catatan = tools.String_Separator_To_String(tmp.Keterangan)
		invent.Jumlah = tools.String_Separator_To_float64(tmp.Jumlah)
		invent.Satuan = tools.String_Separator_To_String(tmp.Satuan)
		invent.Harga = tools.String_Separator_To_Int(tmp.Harga)
		invent.Sub_total = tools.String_Separator_To_Int64(tmp.Sub_total)

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

//Update status Penawaran
func Update_Status_Penawaran(id_proyek string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := " UPDATE penawaran SET status_penawaran=? WHERE id_proyek=?"

	stmt, err := con.Prepare(sqlstatement)

	result, err := stmt.Exec(1, id_proyek)

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

//Update judul Penawaran
func Update_Judul_Penawaran(id_penawaran string, judul string) (tools.Response, error) {
	var res tools.Response

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
func Update_Item_Penawaran(id_penawaran string, id_sub_pekerjaan string, sub_pekerjaan string,
	keterangan string, jumlah float64, satuan string, harga int, sub_total int64) (tools.Response, error) {
	var res tools.Response
	var invent penawaran.Read_Penawaran
	var tmp penawaran.Read_Penawaran_String

	con := db.CreateCon()

	sqlStatement := "SELECT id_sub_pekerjaan, sub_pekerjaan, catatan, jumlah, satuan, harga, sub_total,total FROM penawaran WHERE id_penawaran=? ORDER BY co ASC "

	err := con.QueryRow(sqlStatement, id_penawaran).Scan(&tmp.Id_sub_pekerjaan, &tmp.Sub_pekerjaan, &tmp.Keterangan, &tmp.Jumlah, &tmp.Satuan,
		&tmp.Harga, &tmp.Sub_total, &invent.Total)

	if err != nil {
		return res, err
	}

	invent.Id_sub_pekerjaan = tools.String_Separator_To_String(tmp.Id_sub_pekerjaan)
	invent.Sub_pekerjaan = tools.String_Separator_To_String(tmp.Sub_pekerjaan)
	invent.Catatan = tools.String_Separator_To_String(tmp.Keterangan)
	invent.Jumlah = tools.String_Separator_To_float64(tmp.Jumlah)
	invent.Satuan = tools.String_Separator_To_String(tmp.Satuan)
	invent.Harga = tools.String_Separator_To_Int(tmp.Harga)
	invent.Sub_total = tools.String_Separator_To_Int64(tmp.Sub_pekerjaan)

	sp := ""
	kt := ""
	jmlh := ""
	st := ""
	hg := ""
	sbt := ""
	var tt int64
	tt = 0

	for i := 0; i < len(id_sub_pekerjaan); i++ {
		if invent.Id_sub_pekerjaan[i] == id_sub_pekerjaan {
			invent.Sub_pekerjaan[i] = sub_pekerjaan
			invent.Catatan[i] = keterangan
			invent.Jumlah[i] = jumlah
			invent.Satuan[i] = satuan
			invent.Harga[i] = harga
			invent.Sub_total[i] = sub_total

		}
		tt = tt + invent.Sub_total[i]
		sp = sp + "|" + invent.Sub_pekerjaan[i] + "|"
		kt = kt + "|" + invent.Catatan[i] + "|"
		jmlh = jmlh + "|" + strconv.FormatFloat(invent.Jumlah[i], 'E', -1, 64) + "|"
		st = st + "|" + invent.Satuan[i] + "|"
		hg = hg + "|" + strconv.Itoa(invent.Harga[i]) + "|"
		sbt = sbt + "|" + strconv.FormatInt(invent.Sub_total[i], 64) + "|"
	}

	sqlstatement := "UPDATE penawaran SET id_sub_pekerjaan=?, sub_pekerjaan=?, catatan=?, jumlah=?, satuan=?, harga=?, sub_total=?,total=? WHERE id_penawaran=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(invent.Id_sub_pekerjaan, sp, kt, jmlh, st, hg, sbt, tt, id_penawaran)

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

//Input tambahan sub pekerjaan
func Input_Tambahan_Sub_Pekerjaan(id_proyek string, id_penawaran string, sub_pekerjaan string,
	catatan string, jumlah string, satuan string, harga string, sub_total float64,
	tanggal_pekerjaan_mulai string, durasi int) (tools.Response, error) {
	var res tools.Response

	var rd penawaran.Read_Penawaran_Input_Sub_Pekerjaan

	con := db.CreateCon()

	nm_s := Generate_Id_Sub_Pekerjaan()

	nm_str_s := strconv.Itoa(nm_s)

	id_sub := "|" + "S-" + nm_str_s + "|"

	Sqlstatement := "SELECT id_sub_pekerjaan, sub_pekerjaan,catatan,jumlah,satuan,harga,sub_total,total FROM penawaran WHERE id_penawaran=?"

	_ = con.QueryRow(Sqlstatement, id_penawaran).Scan(&rd.Id_sub_pekerjaan, &rd.Sub_pekerjaan,
		&rd.Catatan, &rd.Jumlah, &rd.Satuan, &rd.Harga, &rd.Sub_total, &rd.Total)

	sqlStatement := "UPDATE penawaran SET id_sub_pekerjaan=?,sub_pekerjaan=?,catatan=?,jumlah=?,satuan=?,harga=?,sub_total=?,total=? WHERE id_penawaran=?"

	rd.Id_sub_pekerjaan = rd.Id_sub_pekerjaan + id_sub
	rd.Sub_pekerjaan = rd.Sub_pekerjaan + "|" + sub_pekerjaan + "|"
	rd.Catatan = rd.Catatan + "|" + catatan + "|"
	rd.Jumlah = rd.Jumlah + "|" + jumlah + "|"
	rd.Satuan = rd.Satuan + "|" + satuan + "|"
	rd.Harga = rd.Harga + "|" + harga + "|"

	sbt_i := int64(math.Round(sub_total*100) / 100)
	sbt := "|" + strconv.FormatInt(sbt_i, 10) + "|"
	rd.Sub_total = rd.Sub_total + "|" + sbt + "|"

	rd.Total = rd.Total + sbt_i

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(rd.Id_sub_pekerjaan, rd.Sub_pekerjaan,
		rd.Catatan, rd.Jumlah, rd.Satuan, rd.Harga, rd.Sub_total,
		rd.Total, id_penawaran)

	nm_p := Generate_Id_Jadwal()

	nm_str_p := strconv.Itoa(nm_p)

	id_jdl := "PJ-" + nm_str_p

	date, _ := time.Parse("02-01-2006", tanggal_pekerjaan_mulai)
	date_sql := date.Format("2006-01-02")

	date_a, _ := time.Parse("2006-01-02", tanggal_pekerjaan_mulai)
	date_awal := date_a.AddDate(0, 0, durasi-1)

	tanggal_Pekerjaan_Selesai := date_awal.Format("2006-01-02")

	sqlStatement = "INSERT INTO penjadwalan (id_sub_pekerjaan,id_penjadwalan,id_proyek,id_penawaran,nama_task,durasi,tanggal_dimulai,tanggal_selesai) values(?,?,?,?,?,?,?,?)"

	stmt, err = con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_sub, id_jdl, id_proyek, id_penawaran, sub_pekerjaan, durasi, date_sql, tanggal_Pekerjaan_Selesai)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//input tambahan pekerjaan tambah
func Input_Tambahan_Pekerjaan_Tambah(id_proyek string, judul string, sub_pekerjaan string,
	catatan string, jumlah string, satuan string, harga string, sub_total string,
	tanggal_mulai string, durasi string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Penawaran()

	nm_str := strconv.Itoa(nm)

	id_penawaran := "P-" + nm_str

	temp := tools.String_Separator_To_String(sub_pekerjaan)

	id_sub_fix := ""

	var id_pj_all []string

	for i := 0; i < len(temp); i++ {
		nm_s := Generate_Id_Sub_Pekerjaan()
		nm_p := Generate_Id_Jadwal()

		nm_str_s := strconv.Itoa(nm_s)
		nm_str_p := strconv.Itoa(nm_p)

		id_sub := "|" + "S-" + nm_str_s + "|"
		id_jdl := "PJ-" + nm_str_p
		id_pj_all = append(id_pj_all, id_jdl)
		id_sub_fix = id_sub_fix + id_sub
	}

	sqlStatement := "INSERT INTO penawaran (id_penawaran,id_proyek,judul,id_sub_pekerjaan,sub_pekerjaan,catatan,jumlah,satuan,harga,total,sub_total,status_penawaran) values(?,?,?,?,?,?,?,?,?,?,?,?)"

	ttl := tools.String_Separator_To_float64(sub_total)
	sbt := ""

	var total int64
	total = 0

	for i := 0; i < len(ttl); i++ {
		sbt_i := int64(math.Round(ttl[i]*100) / 100)
		sbt = sbt + "|" + strconv.FormatInt(sbt_i, 10) + "|"
		total += sbt_i
	}

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_penawaran, id_proyek, judul, id_sub_fix, sub_pekerjaan, catatan, jumlah,
		satuan, harga, sbt, total, 0)

	idsf := tools.String_Separator_To_String(id_sub_fix)
	tm := tools.String_Separator_To_String(tanggal_mulai)
	dur := tools.String_Separator_To_Int(durasi)

	for i := 0; i < len(temp); i++ {

		date, _ := time.Parse("02-01-2006", tm[i])
		date_sql := date.Format("2006-01-02")

		date_a, _ := time.Parse("2006-01-02", tm[i])
		date_awal := date_a.AddDate(0, 0, dur[i]-1)

		tanggal_Pekerjaan_Selesai := date_awal.Format("2006-01-02")

		sqlStatement = "INSERT INTO penjadwalan (id_sub_pekerjaan,id_penjadwalan,id_proyek,id_penawaran,nama_task,tanggal_dimulai,tanggal_selesai,durasi) values(?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(idsf[i], id_pj_all[i], id_proyek, id_penawaran, temp[i],
			date_sql, tanggal_Pekerjaan_Selesai, dur[i])

	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//pilih judul pekerjaan
func Pilih_Judul_Pekerjaan(id_proyek string) (tools.Response, error) {
	var res tools.Response
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

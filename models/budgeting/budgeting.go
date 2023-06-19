package budgeting

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/budgeting"
	"Skripsi/tools"
	"net/http"
	"strconv"
	"time"
)

//Generate Id Penjadwalan
func Generate_Id_Realisasi() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_budgeting FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_budgeting=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input-Realisasi
func Input_Realisasi(id_proyek string, id_sub_pekerjaan string, id_kontrak string,
	perihal_pengeluaran string, tanggal_pembayaran string,
	nominal_pembayaran int64, catatan string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Realisasi()

	nm_str := strconv.Itoa(nm)

	id_real := "BU-" + nm_str

	sqlStatement := "INSERT INTO realisasi (id_realisasi, id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, tanggal_pembayaran, nominal_pembayaran, catatan) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_real, id_proyek, id_sub_pekerjaan, id_kontrak,
		perihal_pengeluaran, date_sql, nominal_pembayaran, catatan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil

}

//Read-Realisasi
func Read_Realisasi(id_proyek string, id_sub_pekerjaan string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []budgeting.Read_Realisasi
	var invent budgeting.Read_Realisasi

	con := db.CreateCon()

	sqlStatement := "SELECT id_realisasi, id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, DATE_FORMAT(tanggal_pembayaran, '%d-%m%-%Y'), nominal_pembayaran, catatan FROM realisasi WHERE id_proyek=? && id_realisasi=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek, id_sub_pekerjaan)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_Realisasi, &invent.Id_Proyek, &invent.Id_Sub_Pekerjaan,
			&invent.Id_Kontrak, &invent.Perihal_Pengeluaran, &invent.Tanggal_Pembayaran,
			&invent.Nominal_Pembayaran, &invent.Catatan)

		if invent.Id_Kontrak != "" {

			sqlStatement = "SELECT nama_vendor FROM kontrak_vendor join vendor on id_master_vendor=id_MV WHERE id_kontrak=?"

			_ = con.QueryRow(sqlStatement, invent.Id_Kontrak).Scan(&invent.Nama_Vendor)

		} else {
			invent.Nama_Vendor = ""
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

//Delete-Realisasi
func Delete_Realisasi(id_realisasi string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "DELETE FROM realisasi WHERE id_realisasi=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_realisasi)

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

//Update-Realisasi
func Update_Realisasi(id_realisasi string, id_kontrak string,
	perihal_pengeluaran string, tanggal_pembayaran string, nominal_pembayaran int64,
	catatan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	sqlstatement := "UPDATE realisasi SET id_kontrak=?,perihal_pengeluaran=?,tanggal_pembayaran=?,nominal_pembayaran=?,catatan=? WHERE id_realisasi=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_kontrak, perihal_pengeluaran, date_sql,
		nominal_pembayaran, catatan, id_realisasi)

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

//Read-Budgeting
func Read_Budgeting(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []budgeting.Read_Budgeting
	var invent budgeting.Read_Budgeting
	var tmp budgeting.Read_Sub_Pekerjaan

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul, id_sub_pekerjaan,sub_pekerjaan, sub_total FROM penawaran WHERE id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		id_sp := ""
		sp := ""
		st := ""

		err = rows.Scan(&invent.Id_penawaran, &invent.Judul, &id_sp, &sp, &st)
		id_sub_pekerjaan := tools.String_Separator_To_String(id_sp)
		sub_pekerjaan := tools.String_Separator_To_String(sp)
		sub_total := tools.String_Separator_To_Int64(st)

		for i := 0; i < len(id_sub_pekerjaan); i++ {
			tmp.Id_sub_pekerjaan = id_sub_pekerjaan[i]
			tmp.Sub_pekerjaan = sub_pekerjaan[i]
			tmp.Biaya_Estimasi = sub_total[i]

			sqlStatement = "SELECT SUM(nominal_pembayaran) FROM realisasi WHERE id_sub_pekerjaan=?"

			_ = con.QueryRow(sqlStatement, id_sub_pekerjaan[i]).Scan(&tmp.Biaya_Realisasi)

			invent.Read_Sub_Pekerjaan = append(invent.Read_Sub_Pekerjaan, tmp)
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

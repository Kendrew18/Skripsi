package budgeting

import (
	"Skripsi/config/db"
	"Skripsi/models/budgeting"
	tools2 "Skripsi/service/tools"
	"fmt"
	"net/http"
	"strconv"
)

//Input-Detail_Budgeting
func Input_Detail_Budgeting(id_proyek string, id_sub_pekerjaan string, id_kontrak string, perihal_pengeluaran string, nominal_pembayaran int64, catatan string) (tools2.Response, error) {

	var res tools2.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM realisasi ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_real := "BU-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO realisasi (co, id_realisasi, id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, nominal_pembayaran, catatan) values(?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_real, id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, nominal_pembayaran, catatan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil

}

//Read-Detail_Budgeting
func Read_Detail_Budgeting(id_proyek string, id_sub_pekerjaan string, id_laporan string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []budgeting.Read_Realisasi
	var invent budgeting.Read_Realisasi

	con := db.CreateCon()

	fmt.Println(id_proyek)
	fmt.Println(id_sub_pekerjaan)

	sqlStatement := "SELECT id_realisasi, id_proyek, id_sub_pekerjaan, id_kontrak, perihal_pengeluaran, nominal_pembayaran, catatan FROM realisasi WHERE id_proyek=? && id_sub_pekerjaan=? && id_laporan=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek, id_sub_pekerjaan, id_laporan)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		fmt.Println("masuk")
		err = rows.Scan(&invent.Id_Realisasi, &invent.Id_Proyek, &invent.Id_Sub_Pekerjaan, &invent.Id_Kontrak, &invent.Perihal_Pengeluaran, &invent.Nominal_Pembayaran, &invent.Catatan)

		fmt.Println(invent)

		if invent.Id_Kontrak != "" {

			sqlStatement = "SELECT nama_vendor FROM kontrak_vendor join vendor on id_master_vendor=id_MV WHERE id_kontrak=?"

			err = con.QueryRow(sqlStatement, invent.Id_Kontrak).Scan(&invent.Nama_Vendor)

			if err != nil {
				return res, err
			}

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

//Delete-Detail_Budgeting
func Delete_Detail_Budgeting(id_realisasi string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	fmt.Println(id_realisasi)

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

//Update-Detail_Budgeting
func Update_Detail_Budgeting(id_realisasi string, id_kontrak string, perihal_pengeluaran string, nominal_pembayaran int64, catatan string) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE realisasi SET id_kontrak=?,perihal_pengeluaran=?,nominal_pembayaran=?,catatan=? WHERE id_realisasi=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_kontrak, perihal_pengeluaran, nominal_pembayaran, catatan, id_realisasi)

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
func Read_Budgeting(id_proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []budgeting.Read_Pengeluaran
	var invent budgeting.Read_Pengeluaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan,DATE_FORMAT(tanggal_laporan, '%d-%m%-%Y') FROM laporan WHERE id_proyek=? ORDER BY co DESC "

	rows, err := con.Query(sqlStatement, id_proyek)

	fmt.Println(id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var rd_sub budgeting.Read_Sub_Pekerjaan
		var arr_rd_sub []budgeting.Read_Sub_Pekerjaan

		err = rows.Scan(&invent.Id_laporan, &invent.Tanggal)

		sqlStatement := "SELECT p.id_sub_pekerjaan,nama_sub_pekerjaan FROM detail_laporan JOIN penjadwalan p on detail_laporan.id_jadwal = p.id_penjadwalan JOIN detail_penawaran dp on p.id_sub_pekerjaan = dp.id_sub_pekerjaan WHERE id_laporan=? ORDER BY detail_laporan.co ASC "

		rows, err := con.Query(sqlStatement, invent.Id_laporan)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&rd_sub.Id_sub_pekerjaan, &rd_sub.Sub_pekerjaan)

			fmt.Println(rd_sub)

			if err != nil {
				return res, err
			}

			arr_rd_sub = append(arr_rd_sub, rd_sub)
		}

		invent.Read_Sub_Pekerjaan = arr_rd_sub

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

//Pilih Kontrak Vendor
func Pilih_Kontrak(id_proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []budgeting.Kontrak_Vendor
	var invent budgeting.Kontrak_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_kontrak,nama_vendor FROM kontrak_vendor JOIN vendor v on kontrak_vendor.id_MV = v.id_master_vendor WHERE id_proyek=? ORDER BY kontrak_vendor.co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_kontak_vendor, &invent.Nama_vendor)
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

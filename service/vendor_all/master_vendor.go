package vendor_all

import (
	"Skripsi/config/db"
	"Skripsi/models/vendor_all"
	"Skripsi/service/tools"
	"fmt"
	"net/http"
	"strconv"
)

//Input_vendor
func Input_Vendor(nama_vendor string, Pekerjaan_Vendor string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM vendor ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_master := "V-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO vendor (co,id_master_vendor, nama_vendor, penkerjaan_vendor) values(?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_master, nama_vendor, Pekerjaan_Vendor)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil

}

//Read_vendor
func Read_Vendor() (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Read_Vendor
	var invent vendor_all.Read_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_master_vendor, nama_vendor, penkerjaan_vendor FROM vendor ORDER BY co ASC "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var arr_tpf []vendor_all.Detail_Read_Vendor_Fix
		var tpf vendor_all.Detail_Read_Vendor_Fix

		err = rows.Scan(&invent.Id_Master_Vendor, &invent.Nama_Vendor, &invent.Pekerjaan_Vendor)

		fmt.Println(invent)

		KTRK := ""

		sqlStatement = "SELECT id_kontrak FROM kontrak_vendor WHERE id_MV=?"

		_ = con.QueryRow(sqlStatement, invent.Id_Master_Vendor).Scan(&KTRK)

		if KTRK != "" {

			sqlStatement = "SELECT COUNT(id_MV) FROM kontrak_vendor WHERE id_MV=? && kontrak_vendor.working_complate=?"

			_ = con.QueryRow(sqlStatement, invent.Id_Master_Vendor, 1).Scan(&invent.Pekerjaan_selesai)

			sqlStatement = "SELECT COUNT(id_MV) FROM kontrak_vendor WHERE id_MV=? && kontrak_vendor.working_complate=?"

			_ = con.QueryRow(sqlStatement, invent.Id_Master_Vendor, 0).Scan(&invent.Pekerjaan_berjalan)

			sqlStatement = "SELECT proyek.id_proyek,p.nama_proyek,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) FROM kontrak_vendor JOIN proyek join proyek p on kontrak_vendor.id_proyek = p.id_proyek WHERE id_MV=? ORDER BY kontrak_vendor.co ASC"

			rows2, err := con.Query(sqlStatement, invent.Id_Master_Vendor)

			defer rows2.Close()

			if err != nil {
				return res, err
			}

			for rows2.Next() {
				durasi := 0

				err = rows2.Scan(&tpf.Id_proyek, &tpf.Nama_proyek, &tpf.Progres, &durasi)
				if err != nil {

					return res, err
				}

				tpf.Progres = (tpf.Progres / durasi) * 100
				arr_tpf = append(arr_tpf, tpf)
			}

			invent.Detail_Vendor = arr_tpf
			if err != nil {
				return res, err
			}

		} else {
			arr_tpf = append(arr_tpf, tpf)
			invent.Detail_Vendor = arr_tpf
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

//Delete_vendor
func Delete_Vendor(id_vendor string) (tools.Response, error) {
	var res tools.Response
	var id_v vendor_all.Id_Vendor

	con := db.CreateCon()

	sqlstatement := "SELECT id_kontrak FROM kontrak_vendor WHERE id_MV=?"

	_ = con.QueryRow(sqlstatement, id_vendor).Scan(&id_v.Id_vendor)

	fmt.Println(id_v.Id_vendor)

	if id_v.Id_vendor == "" {

		sqlstatement = "DELETE FROM vendor WHERE id_master_vendor=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_vendor)

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
		res.Message = "Delete tidak dapat dilakukan"
		res.Data = id_v
	}
	return res, nil
}

//Edit_vendor
func Edit_Vendor(id_master string, nama_vendor string, pekerjaan_vendor string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE vendor SET nama_vendor=?,penkerjaan_vendor=? WHERE id_master_vendor=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_vendor, pekerjaan_vendor, id_master)

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

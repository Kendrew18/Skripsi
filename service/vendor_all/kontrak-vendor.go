package vendor_all

import (
	"Skripsi/config/db"
	"Skripsi/models/vendor_all"
	"Skripsi/service/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Input_Kontrak_Vendor
func Input_Kontrak_Vendor(id_proyek string, id_master_vendor string, total_nilai_kontrak int64, tanggal_dimulai string, tanggal_selesai string, date_pengiriman string, date_dimulai string, date_selesai string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM kontrak_vendor ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_kontrak := "KV-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO kontrak_vendor (co,id_proyek, id_MV, id_kontrak, total_nilai_kontrak, nominal_pembayaran, tanggal_mulai_kontrak, tanggal_berakhir_kontrak, sisa_pembayaran, tanggal_pengiriman, tanggal_pengerjaan_dimulai, tanggal_pengerjaan_berakhir,working_progress,working_complate,kontrak_complate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_dimulai)
	tanggal_dimulai_kontrak_SQL := date.Format("2006-01-02")

	date2, _ := time.Parse("02-01-2006", tanggal_selesai)
	tanggal_berakhir_kontrak_SQL := date2.Format("2006-01-02")

	diff := date2.Sub(date)

	month := int64(diff.Hours()/24/30) + 1

	fmt.Println(month)

	date3, _ := time.Parse("02-01-2006", date_pengiriman)
	Tanggal_Pengiriman_SQL := date3.Format("2006-01-02")

	date4, _ := time.Parse("02-01-2006", date_dimulai)
	Tanggal_Pekerjaan_Dimulai := date4.Format("2006-01-02")

	date5, _ := time.Parse("02-01-2006", date_selesai)
	Tanggal_Pekerjaan_Berakhir := date5.Format("2006-01-02")

	var nominal_pembayaran int64

	nominal_pembayaran = total_nilai_kontrak / month

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_proyek, id_master_vendor, id_kontrak, total_nilai_kontrak, nominal_pembayaran, tanggal_dimulai_kontrak_SQL, tanggal_berakhir_kontrak_SQL, total_nilai_kontrak, Tanggal_Pengiriman_SQL, Tanggal_Pekerjaan_Dimulai, Tanggal_Pekerjaan_Berakhir, 0, 0, 0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Kontrak_Vendor
func Read_Kontrak_Vendor(id_Proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Read_Kontrak_Vendor
	var invent vendor_all.Read_Kontrak_Vendor
	var det_kon vendor_all.Detail_Kontrak_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_kontrak,nama_vendor,penkerjaan_vendor,DATE_FORMAT(tanggal_mulai_kontrak, '%d-%m%-%Y'),DATE_FORMAT(tanggal_berakhir_kontrak, '%d-%m%-%Y'),total_nilai_kontrak,nominal_pembayaran,sisa_pembayaran,DATE_FORMAT(tanggal_pengiriman, '%d-%m%-%Y'),DATE_FORMAT(tanggal_pengerjaan_dimulai, '%d-%m%-%Y'),DATE_FORMAT(tanggal_pengerjaan_berakhir, '%d-%m%-%Y') FROM kontrak_vendor JOIN vendor ON kontrak_vendor.id_MV = vendor.id_master_vendor WHERE id_Proyek=? ORDER BY kontrak_vendor.co ASC "

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_kontak, &invent.Nama_vendor, &invent.Pekerjaan_Vendor,
			&invent.Tanggal_mulai_kontrak, &invent.Tanggal_berakhir_kontrak, &det_kon.Total_nilai_kontrak,
			&det_kon.Nomial_Pembayaran, &det_kon.Sisa_pembayaran, &det_kon.Tanggal_Pengiriman,
			&det_kon.Tanggal_mulai_pengerjaan, &det_kon.Tanggal_berakhir_pengerjaan)
		invent.Detail_Kontrak_Vendor = det_kon
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

//Delete_Kontrak_Vendor
func Delete_Kontrak_Vendor(id_kontrak string) (tools.Response, error) {
	var res tools.Response
	var arrobj []vendor_all.Read_id_pv
	var obj vendor_all.Read_id_pv

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

//Pick_Vendor
func Pick_Vendor() (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Pick_Vendor
	var invent vendor_all.Pick_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_master_vendor, nama_vendor, penkerjaan_vendor FROM vendor ORDER BY co ASC "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&invent.Id_Master_Vendor, &invent.Nama_Vendor,
			&invent.Pekerjaan_Vendor)

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

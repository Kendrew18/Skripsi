package vendor_all

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/vendor_all"
	"Skripsi/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Generate_Id_Kontrak_Vendor
func Generate_Id_Kotrak_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_kv FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_kv=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input_Kontrak_Vendor
func Input_Kontrak_Vendor(id_proyek string, id_master_vendor string,
	total_nilai_kontrak int64, tanggal_dimulai string, tanggal_selesai string,
	date_pengiriman string, date_dimulai string, date_selesai string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Kotrak_Vendor()

	nm_str := strconv.Itoa(nm)

	id_kontrak := "KV-" + nm_str

	sqlStatement := "INSERT INTO kontrak_vendor (id_proyek, id_MV, id_kontrak, total_nilai_kontrak, nominal_pembayaran, tanggal_mulai_kontrak, tanggal_berakhir_kontrak, sisa_pembayaran, tanggal_pengiriman, tanggal_pengerjaan_dimulai, tanggal_pengerjaan_berakhir,working_progress,working_complate,kontrak_complate) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_dimulai)
	date_sql := date.Format("2006-01-02")
	dm := date.Format("200601")

	date2, _ := time.Parse("02-01-2006", tanggal_selesai)
	date_sql2 := date2.Format("2006-01-02")
	dt := date2.Format("200601")

	date3, _ := time.Parse("02-01-2006", date_pengiriman)
	date_sql3 := date3.Format("2006-01-02")

	date4, _ := time.Parse("02-01-2006", date_dimulai)
	date_sql4 := date4.Format("2006-01-02")

	date5, _ := time.Parse("02-01-2006", date_selesai)
	date_sql5 := date5.Format("2006-01-02")

	fmt.Println(dm)
	fmt.Println(dt)

	sqlStatement2 := "SELECT period_diff( " + dt + ", " + dm + ")"

	var temp int64

	temp = 0

	_ = con.QueryRow(sqlStatement2).Scan(&temp)

	var nominal_pembayaran int64

	nominal_pembayaran = 0

	temp += 1

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_proyek, id_master_vendor, id_kontrak, total_nilai_kontrak,
		nominal_pembayaran, date_sql, date_sql2, total_nilai_kontrak, date_sql3,
		date_sql4, date_sql5, 0, 0, 0)

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

	sqlStatement := "SELECT id_kontrak,nama_vendor,penkerjaan_vendor,tanggal_mulai_kontrak,tanggal_berakhir_kontrak,total_nilai_kontrak,nominal_pembayaran,sisa_pembayaran,tanggal_pengiriman,tanggal_pengerjaan_dimulai,tanggal_pengerjaan_berakhir FROM kontrak_vendor JOIN vendor ON kontrak_vendor.id_MV = vendor.id_master_vendor WHERE id_Proyek=? ORDER BY co ASC "

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

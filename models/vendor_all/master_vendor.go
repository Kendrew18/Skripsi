package vendor_all

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/vendor_all"
	"Skripsi/tools"
	"net/http"
	"strconv"
)

//Generate_id_vendor
func Generate_Id_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_MV FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_MV=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input_vendor
func Input_Vendor(nama_vendor string, Pekerjaan_Vendor string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Vendor()

	nm_str := strconv.Itoa(nm)

	id_master := "V-" + nm_str

	sqlStatement := "INSERT INTO vendor (id_master_vendor, nama_vendor, penkerjaan_vendor, proyek_selesai, proyek_berjalan, id_proyek, progres ) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_master, nama_vendor, Pekerjaan_Vendor, 0, 0, "", "")

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
	var tp vendor_all.Detail_Read_Vendor
	var tpf vendor_all.Detail_Read_Vendor_Fix

	con := db.CreateCon()

	sqlStatement := "SELECT id_master_vendor, nama_vendor, penkerjaan_vendor, proyek_selesai, proyek_berjalan, id_proyek, progres FROM vendor ORDER BY co ASC "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_Master_Vendor, &invent.Nama_Vendor, &invent.Pekerjaan_Vendor,
			&invent.Pekerjaan_selesai, &invent.Pekerjaan_berjalan, &tp.Id_proyek, &tp.Nama_proyek, &tp.Progres)

		ip := tools.String_Separator_To_String(tp.Id_proyek)
		np := tools.String_Separator_To_String(tp.Nama_proyek)
		pp := tools.String_Separator_To_Int(tp.Progres)

		for i := 0; i < len(ip); i++ {
			tpf.Id_proyek = ip[i]
			tpf.Nama_proyek = np[i]
			tpf.Progres = pp[i]
			invent.Detail_Vendor = append(invent.Detail_Vendor, tpf)
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

//Delete_vendor
func Delete_Vendor(id_vendor string) (tools.Response, error) {
	var res tools.Response
	var id_v vendor_all.Id_Vendor

	con := db.CreateCon()

	sqlstatement := "SELECT id_master_vendor FROM vendor WHERE proyek_berjalan=? && id_master_vendor=?"

	_ = con.QueryRow(sqlstatement, 0, id_vendor).Scan(&id_v.Id_vendor)

	if id_v.Id_vendor == id_vendor {

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

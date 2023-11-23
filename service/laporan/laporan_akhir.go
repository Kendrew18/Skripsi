package laporan

import (
	"Skripsi/config/db"
	"Skripsi/models/laporan_akhir"
	"Skripsi/service/tools"
	"net/http"
)

func Update_Status(status int, kode_laporan_akhir string) (response tools.Response, err error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "UPDATE laporan_akhir SET status=? WHERE kode_laporan_akhir=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(status, kode_laporan_akhir)

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

func Read_Laporan_Akhir(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []laporan_akhir.Read_Laporan_Akhir
	var invent laporan_akhir.Read_Laporan_Akhir

	con := db.CreateCon()

	sqlStatement := "SELECT kode_laporan_akhir,nama,status FROM laporan_akhir WHERE id_proyek=? ORDER BY co ASC"

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan_akhir, &invent.Nama, &invent.Status)
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

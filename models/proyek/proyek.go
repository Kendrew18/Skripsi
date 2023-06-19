package proyek

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/proyek"
	"Skripsi/tools"
	"net/http"
	"strconv"
)

//Generate-Id-Proyek
func Generate_Id_Proyek() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_proyek=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

//Input-Proyek
func Input_Proyek(id_user string, nama_proyek string, nama_client string, jenis_gedung string,
	alamat string, luas_tanah string, jumlah_lantai string, nama_penanggungjawab_proyek string,
	tanggal_mulai_kerja string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Proyek()

	nm_str := strconv.Itoa(nm)

	kode_proyek := "PRYK-" + nm_str

	jmlt, _ := strconv.Atoi(jumlah_lantai)
	luas_tanah_dbl, _ := strconv.ParseFloat(luas_tanah, 64)

	sqlStatement := "INSERT INTO proyek (id_proyek,id_user,nama_proyek,`nama_client/perusahaan`,alamat,jenis_gedung,luas_tanah,jumlah_lantai,penanggungjawab, tanggal_mulai_kerja) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(kode_proyek, id_user, nama_proyek, nama_client, alamat, jenis_gedung,
		luas_tanah_dbl, jmlt, nama_penanggungjawab_proyek, tanggal_mulai_kerja, 0)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil

}

//Read-Proyek
func Read_Proyek(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []proyek.Read_proyek
	var invent proyek.Read_proyek

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek, nama_proyek, `nama_client/perusahaan`, jenis_gedung, alamat, luas_tanah, jumlah_lantai, penanggungjawab, tanggal_mulai_kerja FROM proyek WHERE status_proyek=? && id_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, 0, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_proyek, &invent.Nama_proyek, &invent.Nama_Client_Perusahaan,
			&invent.Jenis_gedung, &invent.Alamat, &invent.Jumlah_lantai,
			&invent.Luas_tanah, &invent.Penangung_Jawab, &invent.Tanggal_mulai_kerja)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	tmp := ""

	sqlStatement = "SELECT id_penawaran, status_penawaran FROM penawaran WHERE id_proyek=? limit 1"

	_ = con.QueryRow(sqlStatement, id_proyek).Scan(&tmp, &arr_invent[0].Status_penawaran)

	if tmp == "" {
		arr_invent[0].Status_penawaran = 0
	}

	sqlStatement = "SELECT id_penjadwalan, status_urutan FROM penjadwalan WHERE id_proyek=? limit 1"

	_ = con.QueryRow(sqlStatement, id_proyek).Scan(&tmp, &arr_invent[0].Status_penjadwalan)

	if tmp == "" {
		arr_invent[0].Status_penjadwalan = 0
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

//Read-Nama-Proyek
func Read_Nama_Proyek() (tools.Response, error) {
	var res tools.Response
	var arr_invent []proyek.Nama_proyek
	var invent proyek.Nama_proyek

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek,nama_proyek,penanggungjawab FROM proyek where status_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, 0)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_proyek, &invent.Nama_proyek, &invent.Penanggungjawab)
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

//Read-Nama-History-Proyek
func Read_Nama_Proyek_history() (tools.Response, error) {
	var res tools.Response
	var arr_invent []proyek.Nama_proyek
	var invent proyek.Nama_proyek

	con := db.CreateCon()

	sqlStatement := "SELECT nama_proyek FROM proyek WHERE status_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, 1)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Nama_proyek)
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

//Read-History-Proyek
func Read_History() (tools.Response, error) {
	var res tools.Response
	var arr_invent []proyek.Read_proyek
	var invent proyek.Read_proyek

	con := db.CreateCon()

	sqlStatement := "SELECT id_proyek,nama_proyek,jumlah_lantai,luas_tanah,penanggungjawab FROM proyek WHERE status_proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, 1)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_proyek, &invent.Nama_proyek, &invent.Jumlah_lantai, &invent.Luas_tanah, &invent.Penangung_Jawab)
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

//Finish-Proyek
func Finish_Proyek(id_proyek string) (tools.Response, error) {
	var res tools.Response
	con := db.CreateCon()

	sqlstatement := "UPDATE proyek SET status_proyek=? WHERE id_proyek=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_proyek)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

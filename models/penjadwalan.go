package models

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

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

//input_tanggal_mulai
func Input_Tanggal_Mulai(id_proyek string, tanggal string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	date, _ := time.Parse("02-01-2006", tanggal)
	date_sql := date.Format("2006-01-02")

	sqlStatement := "UPDATE proyek SET tanggal_mulai_kerja=? WHERE id_proyek=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(date_sql, id_proyek)

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

//read_tanggal-mulai
func Read_Tanggal_Mulai(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Tanggal_mulai
	var invent str.Tanggal_mulai

	con := db.CreateCon()

	sqlStatement := "SELECT tanggal_mulai_kerja FROM proyek WHERE id_proyek=?"

	err := con.QueryRow(sqlStatement, id_proyek).Scan(&invent.Tanggal_mulai)

	if err != nil {
		return res, err
	}

	arr_invent = append(arr_invent, invent)

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

//input_judul_penawaran
func Read_Judul_Penawaran(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_judul_penawaran
	var invent str.Read_judul_penawaran

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=?"

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

//input_task_penjadwalan
func Input_Task_Penjadwalan(Id_penawaran string, id_proyek string, nama_task string, waktu_optimis float64,
	waktu_pesimis float64, waktu_realistic float64) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Jadwal()

	nm_str := strconv.Itoa(nm)

	id_pjd := "PJD-" + nm_str

	sqlStatement := "INSERT INTO penjadwalan (id_penjadwalan, id_proyek, id_penawaran, nama_task, durasi,dependencies) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	//Three Point Estimating (beta distribution)
	durasi := (waktu_optimis + waktu_pesimis + (4.0 * waktu_realistic)) / 6.0

	durasi_int := int(durasi)
	durasi_double := float64(durasi_int)
	real_durasi := 0

	if durasi_double == durasi {
		real_durasi = durasi_int
	} else {
		real_durasi = durasi_int + 1
	}
	_, err = stmt.Exec(id_pjd, id_proyek, Id_penawaran, nama_task, real_durasi, "")

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//read_task
func Read_Task(id_proyek string, id_penawaran string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Task
	var invent str.Read_Task

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,nama_task,durasi,dependencies FROM penjadwalan WHERE id_proyek=? && id_penawaran=?"

	rows, err := con.Query(sqlStatement, id_proyek, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penjadwalan, &invent.Nama_Task, &invent.Durasi_Task,
			&invent.Dependentcies)
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

//input_depedenytcies
func Input_Dependentcies(id_jadwal string, dep string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE penjadwalan SET dependencies=? WHERE id_penjadwalan=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(dep, id_jadwal)

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

//generate_jadwal
func Generate_Jadwal(id_proyek string) (tools.Response, error) {
	//urutno
	var res tools.Response
	var arr_invent []str.Gene_JDL
	var invent str.Gene_JDL

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,durasi,dependencies FROM penjadwalan WHERE id_Proyek=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id, &invent.Durasi, &invent.Dependentcies)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	var arr_invent_fn []str.Gene_JDL

	var arr_invent_blm []str.Gene_JDL

	urt := 1

	for i := 0; i < len(arr_invent); i++ {
		if arr_invent[i].Dependentcies == "" {
			arr_invent[i].Status_urutan = urt
			arr_invent_fn = append(arr_invent_fn, arr_invent[i])
		} else {
			arr_invent_blm = append(arr_invent_blm, arr_invent[i])
		}
	}

	arr_invent = arr_invent_blm

	urt++

	co := 0
	i := -1
	var arr_index []int

	for 0 < len(arr_invent) {
		i++

		arr_dep := tools.String_Separator_To_String(arr_invent[i].Dependentcies)
		co_dep := len(arr_dep)

		for j := 0; j < len(arr_invent_fn); j++ {
			for k := 0; k < len(arr_dep); k++ {
				if arr_invent_fn[j].Id == arr_dep[k] {
					co++
				}
			}
		}

		if co == co_dep {
			arr_index = append(arr_index, i)
		}

		if i == len(arr_invent)-1 {
			arr_invent_blm = []str.Gene_JDL{}
			for j := 0; j < len(arr_invent); j++ {
				for k := 0; k < len(arr_index); k++ {
					if arr_index[k] == j {
						arr_invent[j].Status_urutan = urt
						arr_invent_fn = append(arr_invent_fn, arr_invent[j])
					}
				}
			}
			test := 0
			for j := 0; j < len(arr_invent); j++ {
				for k := 0; k < len(arr_index); k++ {
					if arr_index[k] == j {
						test++
					}
				}
				if test == 0 {
					arr_invent_blm = append(arr_invent_blm, arr_invent[j])
				}
				test = 0
			}

			arr_invent = arr_invent_blm

			i = -1
			urt++
			arr_index = []int{}

		}

		co = 0

	}

	//CPM
	//ES EF
	for i := 0; i < len(arr_invent_fn); i++ {
		if arr_invent_fn[i].Status_urutan == 1 {
			arr_invent_fn[i].Es = 0
			arr_invent_fn[i].Ef = arr_invent_fn[i].Durasi
		}
	}

	x := 1

	for x < urt {
		x++
		for i := 0; i < len(arr_invent_fn); i++ {
			if arr_invent_fn[i].Status_urutan == x {
				dep := tools.String_Separator_To_String(arr_invent_fn[i].Dependentcies)
				max := 0
				for k := 0; k < len(arr_invent_fn); k++ {
					for j := 0; j < len(dep); j++ {
						if arr_invent_fn[k].Id == dep[j] {
							if max < arr_invent_fn[k].Ef {
								max = arr_invent_fn[k].Ef
							}
						}
					}
				}
				arr_invent_fn[i].Es = max
				arr_invent_fn[i].Ef = max + arr_invent_fn[i].Durasi

			}
		}
	}

	for i := 0; i < len(arr_invent_fn); i++ {
		if arr_invent_fn[i].Status_urutan == 1 {
			arr_invent_fn[i].Es = 0
			arr_invent_fn[i].Ef = arr_invent_fn[i].Durasi
		}
	}

	//LS LF
	for i := 0; i < len(arr_invent_fn); i++ {
		if arr_invent_fn[i].Status_urutan == urt-1 {
			arr_invent_fn[i].Lf = arr_invent_fn[i].Ef
			arr_invent_fn[i].Ls = arr_invent_fn[i].Lf - arr_invent_fn[i].Durasi
		}
	}

	y := urt

	for y > 0 {
		y--
		for i := 0; i < len(arr_invent_fn); i++ {
			if arr_invent_fn[i].Status_urutan == y {

				if arr_invent_fn[i].Lf == 0 {
					arr_invent_fn[i].Lf = arr_invent_fn[i].Ef
					arr_invent_fn[i].Ls = arr_invent_fn[i].Ef - arr_invent_fn[i].Durasi
				}

				dep := tools.String_Separator_To_String(arr_invent_fn[i].Dependentcies)

				for k := 0; k < len(arr_invent_fn); k++ {
					for j := 0; j < len(dep); j++ {
						if arr_invent_fn[k].Id == dep[j] {
							if arr_invent_fn[k].Lf == 0 {
								arr_invent_fn[k].Lf = arr_invent_fn[i].Ls
								arr_invent_fn[k].Ls = arr_invent_fn[k].Lf - arr_invent_fn[k].Durasi
							} else if arr_invent_fn[k].Lf > arr_invent_fn[i].Ls {
								arr_invent_fn[k].Lf = arr_invent_fn[i].Ls
								arr_invent_fn[k].Ls = arr_invent_fn[k].Lf - arr_invent_fn[k].Durasi
							}

						}
					}
				}
			}
		}
	}

	var tgl str.Tanggal_mulai

	sqlStatement = "SELECT tanggal_mulai_kerja FROM proyek WHERE id_proyek=?"
	_ = con.QueryRow(sqlStatement, id_proyek).Scan(&tgl.Tanggal_mulai)

	//float
	for i := 0; i < len(arr_invent_fn); i++ {
		arr_invent_fn[i].Tf = arr_invent_fn[i].Lf - arr_invent_fn[i].Ef
		arr_invent_fn[i].Ff = arr_invent_fn[i].Ls - arr_invent_fn[i].Es

		date_a, _ := time.Parse("2006-01-02", tgl.Tanggal_mulai)
		date_awal := date_a.AddDate(0, 0, arr_invent_fn[i].Es)

		arr_invent_fn[i].Tanggal_mulai = date_awal.Format("2006-01-02")

		date_b, _ := time.Parse("2006-01-02", tgl.Tanggal_mulai)
		date_akhir := date_b.AddDate(0, 0, arr_invent_fn[i].Ef-1)

		arr_invent_fn[i].Tanggal_berakhir = date_akhir.Format("2006-01-02")

	}

	fmt.Println(arr_invent_fn)

	//tanggal

	for j := 0; j < len(arr_invent_fn); j++ {

		sqlStatement = "UPDATE penjadwalan SET es=?,ls=?,ef=?,lf=?,tf=?,ff=?,tanggal_dimulai=?,tanggal_selesai=?,status_urutan=? WHERE id_penjadwalan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(arr_invent_fn[j].Es, arr_invent_fn[j].Ls, arr_invent_fn[j].Ef, arr_invent_fn[j].Lf,
			arr_invent_fn[j].Tf, arr_invent_fn[j].Ff, arr_invent_fn[j].Tanggal_mulai,
			arr_invent_fn[j].Tanggal_berakhir, arr_invent_fn[j].Status_urutan, arr_invent_fn[j].Id)

		if err != nil {
			return res, err
		}

	}

	if arr_invent_fn == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent_fn
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent_fn
	}

	return res, nil

}

//Read_jadwal
func Read_Jadwal(id_proyek string, id_penawaran string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Read_Task
	var invent str.Read_Task

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,nama_task,tanggal_dimulai,tanggal_selesai FROM penjadwalan WHERE id_proyek=? && id_penawaran=?"

	rows, err := con.Query(sqlStatement, id_proyek, id_penawaran)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penjadwalan, &invent.Nama_Task, &invent.Durasi_Task,
			&invent.Dependentcies)
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

package jadwal

import (
	"Skripsi/db"
	"Skripsi/struct_all/jadwal"
	"Skripsi/tools"
	"fmt"
	"net/http"
	"time"
)

//input_Durasi_task(done)
func Input_Durasi_task(Id_penjadwalan string, waktu_optimis float64,
	waktu_pesimis float64, waktu_realistic float64) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE penjadwalan SET durasi=?,dependencies=? WHERE id_penjadwalan=?"

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
	_, err = stmt.Exec(real_durasi, "", Id_penjadwalan)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//read_task (done)
func Read_Task(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []jadwal.Read_Task
	var invent jadwal.Read_Task
	var in jadwal.Sub_Task

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=?"

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penawaran, &invent.Judul_penawaran)

		sqlStatement = "SELECT id_penjadwalan,nama_task,durasi,dependencies FROM penjadwalan WHERE id_proyek=? && id_penawaran=?"

		rows, err = con.Query(sqlStatement, id_proyek, invent.Id_penawaran)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&in.Id_penjadwalan, &in.Nama_Task, &in.Durasi_Task,
				&in.Dependentcies)
			if err != nil {
				return res, err
			}
			invent.Sub_task = append(invent.Sub_task, in)
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

//read depedentcies (done)
func Read_dep(id_proyek string, id_penjadwalan string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []jadwal.Read_dep
	var invent jadwal.Read_dep

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,nama_task FROM penjadwalan WHERE id_proyek=? && id_penjadwalan!=?"

	rows, err := con.Query(sqlStatement, id_proyek, id_penjadwalan)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penjadwalan, &invent.Nama_penjadwalan)
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

//input_depedentcies(done)
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

//generate_jadwal(done)
func Generate_Jadwal(id_proyek string) (tools.Response, error) {
	//urutno
	var res tools.Response
	var arr_invent []jadwal.Gene_JDL
	var invent jadwal.Gene_JDL

	con := db.CreateCon()

	sqlStatement := "SELECT id_penjadwalan,durasi,dependencies FROM penjadwalan WHERE id_proyek=? && penjadwalan.status_urutan!=-1 ORDER BY co ASC "

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

	fmt.Println(arr_invent)

	var arr_invent_fn []jadwal.Gene_JDL

	var arr_invent_blm []jadwal.Gene_JDL

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
	fmt.Println(arr_invent)
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
			arr_invent_blm = []jadwal.Gene_JDL{}
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

	var tgl jadwal.Tanggal_mulai

	sqlStatement = "SELECT tanggal_mulai_kerja FROM proyek WHERE id_proyek=?"

	_ = con.QueryRow(sqlStatement, id_proyek).Scan(&tgl.Tanggal_mulai)

	//float dan tanggal
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

	for j := 0; j < len(arr_invent_fn); j++ {

		sqlStatement = "UPDATE penjadwalan SET es=?,ls=?,ef=?,lf=?,tf=?,ff=?,tanggal_dimulai=?,tanggal_selesai=?,status_urutan=?,progress=? WHERE id_penjadwalan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(arr_invent_fn[j].Es, arr_invent_fn[j].Ls, arr_invent_fn[j].Ef, arr_invent_fn[j].Lf,
			arr_invent_fn[j].Tf, arr_invent_fn[j].Ff, arr_invent_fn[j].Tanggal_mulai,
			arr_invent_fn[j].Tanggal_berakhir, arr_invent_fn[j].Status_urutan, arr_invent_fn[j].Id, 0)

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

//Read_jadwal (done)
func Read_Jadwal(id_proyek string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []jadwal.Read_Task_Jadwal
	var invent jadwal.Read_Task_Jadwal
	var in jadwal.Sub_Task_Jadwal

	con := db.CreateCon()

	sqlStatement := "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=?"

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_penawaran, &invent.Judul_penawaran)

		sqlStatement = "SELECT id_penjadwalan,nama_task,DATE_FORMAT(tanggal_dimulai, '%d-%m%-%Y'),DATE_FORMAT(tanggal_selesai, '%d-%m%-%Y') FROM penjadwalan WHERE id_proyek=? && id_penawaran=?"

		rows, err = con.Query(sqlStatement, id_proyek, invent.Id_penawaran)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&in.Id_penjadwalan, &in.Nama_Task, &in.Tanggal_Mulai,
				&in.Tanggal_Selesai)
			if err != nil {
				return res, err
			}
			invent.Sub_task = append(invent.Sub_task, in)
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

//edit jadwal tanggal mulai dan durasinya (done)
func Edit_Dur_Tgl(id_penjadwalan string, tanggal_mulai string, durasi int) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := " UPDATE penjadwalan SET tanggal_dimulai=?,durasi=?,tanggal_selesai=? WHERE id_penjadwalan=?"

	date, _ := time.Parse("02-01-2006", tanggal_mulai)
	date_sql := date.Format("2006-01-02")

	date_a, _ := time.Parse("2006-01-02", tanggal_mulai)
	date_awal := date_a.AddDate(0, 0, durasi-1)

	tanggal_Pekerjaan_Selesai := date_awal.Format("2006-01-02")

	stmt, err := con.Prepare(sqlstatement)

	result, err := stmt.Exec(date_sql, durasi, tanggal_Pekerjaan_Selesai, id_penjadwalan)

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

//see_calender (done)
func See_Calender(id_proyek string, status_user int) (tools.Response, error) {
	var res tools.Response
	var arr_invent []jadwal.See_Calender
	var invent jadwal.See_Calender

	var arr_vendor []jadwal.See_Calender_Vendor
	var vendor jadwal.See_Calender_Vendor

	con := db.CreateCon()

	//Penjadwalan
	sqlStatement := "SELECT nama_task,DATE_FORMAT(tanggal_dimulai, '%d-%m%-%Y'),DATE_FORMAT(tanggal_selesai, '%d-%m%-%Y') FROM penjadwalan WHERE id_proyek=?"

	rows, err := con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Judul, &invent.Tanggal_mulai, &invent.Tanggal_selesai)
		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	if status_user == 1 || status_user == 3 {
		//Pembayaran Vendor
		sqlStatement = "SELECT id_kontrak,nama_vendor,nominal_pembayaran,tanggal_mulai_kontrak,tanggal_berakhir_kontrak FROM kontrak_vendor JOIN vendor ON id_MV=id_master_vendor WHERE kontrak_vendor.id_proyek=?"

		rows, err = con.Query(sqlStatement, id_proyek)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&vendor.Id_kontrak, &vendor.Nama_vendor, &vendor.Nominal_Pembayaran,
				&vendor.Tanggal_mulai, &vendor.Tanggal_selesai)
			if err != nil {
				return res, err
			}
			arr_vendor = append(arr_vendor, vendor)
		}

		for i := 0; i < len(arr_vendor); i++ {
			judul := arr_vendor[i].Id_kontrak + " | " + arr_vendor[i].Nama_vendor + " | " + arr_vendor[i].Nominal_Pembayaran
			x := 0

			date, _ := time.Parse("2006-01-02", arr_vendor[i].Tanggal_mulai)
			dm := date.Format("200601")

			date2, _ := time.Parse("2006-01-02", arr_vendor[i].Tanggal_selesai)
			dt := date2.Format("200601")

			sqlStatement2 := "SELECT period_diff( " + dt + ", " + dm + ")"

			var temp int

			temp = 0

			_ = con.QueryRow(sqlStatement2).Scan(&temp)

			temp += 1

			for x < temp {
				date_awal := date.AddDate(0, x, 0)
				date_sql := date_awal.Format("02-01-2006")

				invent.Judul = judul
				invent.Tanggal_mulai = date_sql
				invent.Tanggal_selesai = date_sql

				arr_invent = append(arr_invent, invent)
				x++
			}
		}
	}

	//kerja_vendor
	sqlStatement = "SELECT id_kontrak,nama_vendor,tanggal_pengerjaan_dimulai,tanggal_pengerjaan_berakhir FROM kontrak_vendor JOIN vendor ON id_MV=id_master_vendor WHERE kontrak_vendor.id_proyek=?"

	rows, err = con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&vendor.Id_kontrak, &vendor.Nama_vendor,
			&vendor.Tanggal_mulai, &vendor.Tanggal_selesai)
		if err != nil {
			return res, err
		}
		arr_vendor = append(arr_vendor, vendor)
	}

	for i := 0; i < len(arr_vendor); i++ {
		judul := arr_vendor[i].Id_kontrak + " | " + arr_vendor[i].Nama_vendor

		date, _ := time.Parse("2006-01-02", arr_vendor[i].Tanggal_mulai)
		date_sql := date.Format("02-01-2006")

		date2, _ := time.Parse("2006-01-02", arr_vendor[i].Tanggal_selesai)
		date_sql2 := date2.Format("02-01-2006")

		invent.Judul = judul
		invent.Tanggal_mulai = date_sql
		invent.Tanggal_selesai = date_sql2
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

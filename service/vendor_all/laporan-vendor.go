package vendor_all

import (
	"Skripsi/config/db"
	"Skripsi/models/jadwal"
	"Skripsi/models/vendor_all"
	tools2 "Skripsi/service/tools"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//input-laporan-Vendor (done)V
func Input_Laporan_Vendor(id_proyek string, laporan string, tanggal_laporan string, id_kontrak string, check_box string) (tools2.Response, error) {
	var res tools2.Response
	var RP vendor_all.Progress_Vendor

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM laporan_vendor ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_LV := "LV-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO laporan_vendor (co,id_proyek,id_laporan_vendor,laporan,tanggal_laporan,status_laporan) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_proyek, id_LV, laporan, date_sql, 0)

	if err != nil {
		return res, err
	}

	ip := tools2.String_Separator_To_String(id_kontrak)
	ck := tools2.String_Separator_To_Int(check_box)

	for i := 0; i < len(ip); i++ {
		fmt.Println("masuk")
		progress := 0
		if ck[i] == 1 {
			sqlstatemen_jdl := "SELECT id_kontrak,working_progress,working_complate,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) FROM kontrak_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess, &RP.Working_Complate, &RP.Durasi)

			if RP.Durasi == RP.Working_Progess+1 {
				RP.Working_Progess++
				RP.Working_Complate = 1
			} else {
				RP.Working_Progess = RP.Working_Progess + 1
			}

			sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				fmt.Println("masuk2")
				return res, err
			}

			_, err = stmt.Exec(RP.Working_Progess, RP.Working_Complate, ip[i])

			progress_float := float64(RP.Working_Progess)
			durasi_float := float64(RP.Durasi)

			progress = int(math.Round((progress_float / durasi_float) * 100))
			fmt.Println(RP.Durasi)
			fmt.Println(progress)

		} else if ck[i] == 2 {

			sqlStatement = "UPDATE kontrak_vendor SET working_complate=? WHERE id_kontrak=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(1, ip[i])

			progress = 100
		}
		nm_str_DLPV := 0

		Sqlstatement := "SELECT co FROM detail_laporan_vendor ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DLPV)

		nm_str_DLPV = nm_str_DLPV + 1

		id_detail_laporan_vendor := "DLPV-" + strconv.Itoa(nm_str_DLPV)

		sqlStatement := "INSERT INTO detail_laporan_vendor (co, id_detail_laporan_vendor, id_laporan_vendor, id_kontrak, check_box, progress) values(?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str_DLPV, id_detail_laporan_vendor, id_LV, ip[i], ck[i], progress)

		if err != nil {
			return res, err
		}

	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//read-laporan-Vendor(done)V
func Read_Laporan_Vendor(id_Proyek string) (tools2.Response, error) {
	var res tools2.Response
	var arr_invent []vendor_all.Read_Laporan_Vendor
	var invent vendor_all.Read_Laporan_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan_vendor, laporan, DATE_FORMAT(tanggal_laporan, '%d-%m%-%Y'),status_laporan FROM laporan_vendor WHERE laporan_vendor.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_laporan_vendor, &invent.Laporan, &invent.Tanggal_laporan, &invent.Status_laporan)

		if err != nil {
			return res, err
		}
		arr_invent = append(arr_invent, invent)
	}

	for i := 0; i < len(arr_invent); i++ {
		var RlV vendor_all.Id_Kontrak
		//Read Detail Laporan

		sqlStatement := "SELECT id_kontrak,progress FROM detail_laporan_vendor WHERE id_laporan_vendor=?"

		rows, err := con.Query(sqlStatement, arr_invent[i].Id_laporan_vendor)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			var dl vendor_all.Detail_Laporan_Vendor

			err = rows.Scan(&RlV.Id_Kontrak, &dl.Progress)

			if err != nil {
				return res, err
			}

			durasi := 0
			complate := 0

			sqlstatemen_jdl := "SELECT id_kontrak,nama_vendor,penkerjaan_vendor,datediff(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor JOIN vendor v on kontrak_vendor.id_MV=v.id_master_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlstatemen_jdl, RlV.Id_Kontrak).Scan(&dl.Id_kontrak, &dl.Nama_vendor, &dl.Pekerjaan_vendor, &durasi, &complate)

			if complate == 1 && durasi != dl.Progress {
				dl.Progress = durasi
			}

			arr_invent[i].Detail_Laporan_Vendor = append(arr_invent[i].Detail_Laporan_Vendor, dl)

		}
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

//Update-laporan-Vendor (done)V
func Update_Laporan_Vendor(id_laporan_vendor string, laporan string, id_kontrak string, check_box string) (tools2.Response, error) {

	var res tools2.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&st.Status)

	cont := 0

	sqls := "SELECT count(id_laporan_vendor) FROM detail_laporan_vendor WHERE id_laporan_vendor=?"
	_ = con.QueryRow(sqls, id_laporan_vendor).Scan(&cont)

	fmt.Println(cont)

	if cont == 0 && st.Status == 0 {
		var RP vendor_all.Progress_Vendor

		ip := tools2.String_Separator_To_String(id_kontrak)
		ck := tools2.String_Separator_To_Int(check_box)

		for i := 0; i < len(ip); i++ {
			fmt.Println("masuk")
			progress := 0
			if ck[i] == 1 {
				sqlstatemen_jdl := "SELECT id_kontrak,working_progress,working_complate,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) FROM kontrak_vendor WHERE id_kontrak=?"

				_ = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess, &RP.Working_Complate, &RP.Durasi)

				if RP.Durasi == RP.Working_Progess+1 {
					RP.Working_Progess++
					RP.Working_Complate = 1
				} else {
					RP.Working_Progess = RP.Working_Progess + 1
				}

				sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					fmt.Println("masuk2")
					return res, err
				}

				_, err = stmt.Exec(RP.Working_Progess, RP.Working_Complate, ip[i])

				progress_float := float64(RP.Working_Progess)
				durasi_float := float64(RP.Durasi)

				progress = int(math.Round((progress_float / durasi_float) * 100))
				fmt.Println(RP.Durasi)
				fmt.Println(progress)

			} else if ck[i] == 2 {

				sqlStatement = "UPDATE kontrak_vendor SET working_complate=? WHERE id_kontrak=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(1, ip[i])

				progress = 100
			}
			nm_str_DLPV := 0

			Sqlstatement := "SELECT co FROM detail_laporan_vendor ORDER BY co DESC Limit 1"

			_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DLPV)

			nm_str_DLPV = nm_str_DLPV + 1

			id_detail_laporan_vendor := "DLPV-" + strconv.Itoa(nm_str_DLPV)

			sqlStatement := "INSERT INTO detail_laporan_vendor (co, id_detail_laporan_vendor, id_laporan_vendor, id_kontrak, check_box, progress) values(?,?,?,?,?,?)"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(nm_str_DLPV, id_detail_laporan_vendor, id_laporan_vendor, ip[i], ck[i], progress)

			if err != nil {
				return res, err
			}

		}

		sqlStatement = "UPDATE laporan_vendor SET laporan=? WHERE id_laporan_vendor=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, id_laporan_vendor)

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

		stmt.Close()

	} else if cont != 0 && st.Status == 0 {

		var read_dt_lp vendor_all.Detail_Laporan_Vendor_Update
		var arr_read_dt_lp []vendor_all.Detail_Laporan_Vendor_Update
		var rp vendor_all.Progress_Vendor
		var RP vendor_all.Progress_Vendor

		ip := tools2.String_Separator_To_String(id_kontrak)
		ck := tools2.String_Separator_To_Int(check_box)

		sqlStatement := "SELECT co, id_detail_laporan_vendor,id_kontrak,check_box FROM detail_laporan_vendor"

		q1 := " WHERE id_kontrak NOT IN ("

		for i := 0; i < len(ip); i++ {
			if i == len(ip)-1 {
				q1 = q1 + "'" + ip[i] + "') && id_laporan_vendor = " + "'" + id_laporan_vendor + "' "
			} else {
				q1 = q1 + "'" + ip[i] + "' , "
			}
		}

		sqlStatement = sqlStatement + q1

		rows, err := con.Query(sqlStatement)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&read_dt_lp.Co, &read_dt_lp.Id_Detail_Laporan_Vendor, &read_dt_lp.Id_Kontrak_Vendor, &read_dt_lp.Check_Box)

			if err != nil {
				return res, err
			}
			arr_read_dt_lp = append(arr_read_dt_lp, read_dt_lp)
		}

		fmt.Println(arr_read_dt_lp, len(arr_read_dt_lp))

		for i := 0; i < len(arr_read_dt_lp); i++ {
			//Update Penjadwalan
			fmt.Println("Masuk")

			sqlStatement = "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

			_ = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Id_Kontrak_Vendor).Scan(&rp.Id_kontrak, &rp.Working_Progess, &rp.Durasi, &rp.Working_Complate)

			if rp.Working_Complate == 1 && arr_read_dt_lp[i].Check_Box == 1 {
				rp.Working_Complate = 0
				rp.Working_Progess--
			} else if rp.Working_Complate == 1 {
				rp.Working_Complate = 0
			} else if rp.Working_Complate == 0 {
				rp.Working_Progess--
			}

			sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(rp.Working_Progess, rp.Working_Complate, rp.Id_kontrak)

			if err != nil {
				return res, err
			}

			//update penjadwalan
			var ARR_Id_Detail_Laporan []vendor_all.Id_Detail_Laporan_Vendor
			var Id_Detail_Laporan vendor_all.Id_Detail_Laporan_Vendor

			progres_lama := 0

			sqlStatement = "SELECT COUNT(id_detail_laporan_vendor) FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co < ? && id_kontrak=?"

			err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor).Scan(&progres_lama)

			sqlStatement = "SELECT id_detail_laporan_vendor FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co > ? && id_kontrak=? ORDER BY lv.co ASC"

			rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor)

			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&Id_Detail_Laporan.Id_Detail_Laporan_Vendor)

				if err != nil {
					return res, err
				}
				ARR_Id_Detail_Laporan = append(ARR_Id_Detail_Laporan, Id_Detail_Laporan)
			}

			for x := 0; x < len(ARR_Id_Detail_Laporan); x++ {
				progres_lama++
				progress_float := float64(progres_lama)
				durasi_float := float64(rp.Durasi)

				progress := int(math.Round((progress_float / durasi_float) * 100))

				sqlStatement = "UPDATE detail_laporan SET progress=? WHERE id_detail_laporan=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_Detail_Laporan_Vendor)

				if err != nil {
					return res, err
				}
			}

			// menghilangkan detail laporan yang gak perlu --

			q2 := " WHERE id_detail_laporan_vendor IN ("

			for j := 0; j < len(arr_read_dt_lp); j++ {
				if j == len(arr_read_dt_lp)-1 {
					q2 = q2 + "'" + arr_read_dt_lp[i].Id_Detail_Laporan_Vendor + "') && id_laporan_vendor = " + "'" + id_laporan_vendor + "' "
				} else {
					q2 = q2 + "'" + arr_read_dt_lp[i].Id_Detail_Laporan_Vendor + "' , "
				}
			}

			sqlstatement := "DELETE FROM detail_laporan_vendor"

			sqlstatement = sqlstatement + q2

			fmt.Println(sqlstatement)

			stmt, err = con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			result, err := stmt.Exec()

			if err != nil {
				return res, err
			}

			_, err = result.RowsAffected()

			if err != nil {
				return res, err
			}
		}

		for i := 0; i < len(ip); i++ {
			id_detail_laporan := ""
			check_box2 := 0

			sqlStatement = "SELECT id_detail_laporan_vendor, check_box FROM detail_laporan_vendor WHERE id_laporan_vendor=? && id_kontrak=?"

			_ = con.QueryRow(sqlStatement, id_laporan_vendor, ip[i]).Scan(&id_detail_laporan, &check_box2)

			fmt.Println(id_detail_laporan)

			if id_detail_laporan == "" {
				progress := 0
				if ck[i] == 1 {
					sqlstatemen_jdl := "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

					err = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess, &RP.Durasi, &RP.Working_Complate)

					if err != nil {
						fmt.Println("masuk 3")
						return res, err
					}

					if RP.Durasi == RP.Working_Progess+1 {
						RP.Working_Progess++
						RP.Working_Complate = 1
					} else {
						RP.Working_Progess = RP.Working_Progess + 1
					}

					sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(RP.Working_Progess, RP.Working_Complate, ip[i])

					if err != nil {
						return res, err
					}

					var ARR_Id_Detail_Laporan []vendor_all.Id_Detail_Laporan_Vendor
					var Id_Detail_Laporan vendor_all.Id_Detail_Laporan_Vendor

					progres_lama := 0

					sqlStatement = "SELECT COUNT(id_detail_laporan_vendor) FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co < ? && id_kontrak=?"

					err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor).Scan(&progres_lama)

					progres_lama = progres_lama + 1

					progress_float := float64(RP.Working_Progess)
					durasi_float := float64(RP.Durasi)

					progress = int(math.Round((progress_float / durasi_float) * 100))

					sqlStatement = "SELECT id_detail_laporan_vendor FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co > ? && id_kontrak=? ORDER BY lv.co ASC"

					rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor)

					defer rows.Close()

					for rows.Next() {
						err = rows.Scan(&Id_Detail_Laporan.Id_Detail_Laporan_Vendor)

						if err != nil {
							return res, err
						}
						ARR_Id_Detail_Laporan = append(ARR_Id_Detail_Laporan, Id_Detail_Laporan)
					}

					for x := 0; x < len(ARR_Id_Detail_Laporan); x++ {
						progres_lama++
						progress_float := float64(progres_lama)
						durasi_float := float64(rp.Durasi)

						progress := int(math.Round((progress_float / durasi_float) * 100))

						sqlStatement = "UPDATE detail_laporan SET progress=? WHERE id_detail_laporan=?"

						stmt, err := con.Prepare(sqlStatement)

						if err != nil {
							return res, err
						}

						_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_Detail_Laporan_Vendor)

						if err != nil {
							return res, err
						}
					}

				} else if ck[i] == 2 {

					sqlStatement = "UPDATE penjadwalan SET complate=? WHERE id_penjadwalan=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(1, ip[i])

					if err != nil {
						return res, err
					}

					progress = 100
				}

				nm_str_DLP := 0

				Sqlstatement := "SELECT co FROM detail_laporan_vendor ORDER BY co DESC Limit 1"

				err = con.QueryRow(Sqlstatement).Scan(&nm_str_DLP)

				fmt.Println(nm_str_DLP)

				if err != nil {
					nm_str_DLP = 0
				}

				nm_str_DLP = nm_str_DLP + 1

				id_detail_laporan_vendor := "DLPV-" + strconv.Itoa(nm_str_DLP)

				sqlStatement := "INSERT INTO detail_laporan_vendor (co, id_detail_laporan_vendor, id_laporan_vendor, id_kontrak, check_box, progress) values(?,?,?,?,?,?)"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(nm_str_DLP, id_detail_laporan_vendor, id_laporan_vendor, ip[i], ck[i], progress)

				if err != nil {
					return res, err
				}

			} else if check_box2 != ck[i] {
				progress := 0
				if ck[i] == 2 {

					sqlStatement = "UPDATE kontrak_vendor SET working_complate=? WHERE id_kontrak=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(1, ip[i])

					progress = 100

				} else if ck[i] == 1 {
					prg_lama := 0
					durasi := 0
					sqlStatement = "SELECT DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_progress FROM kontrak_vendor WHERE id_kontrak=?"

					err = con.QueryRow(sqlStatement, ip[i]).Scan(&durasi, &prg_lama)

					if err != nil {
						return res, err
					}

					sqlStatement = "UPDATE penjadwalan SET complate=?,progress=? WHERE id_penjadwalan=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					prg_lama++

					cplt := 0

					if prg_lama == durasi {
						cplt = 1
					}

					_, err = stmt.Exec(cplt, prg_lama, ip[i])

					progress_float := float64(prg_lama)
					durasi_float := float64(durasi)

					progress = int(math.Round((progress_float / durasi_float) * 100))
				}

				sqlStatement = "UPDATE detail_laporan_vendor SET progress=? WHERE id_detail_laporan_vendor=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(progress, id_detail_laporan)

				if err != nil {
					return res, err
				}

			}
		}

		sqlStatement = "UPDATE laporan_vendor SET laporan=? WHERE id_laporan_vendor=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, id_laporan_vendor)

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
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
	}

	return res, nil
}

//Update-Status-laporan-Vendor (done)V
func Update_Status_Laporan_Vendor(id_laporan_vendor string) (tools2.Response, error) {
	var res tools2.Response
	var RP vendor_all.Progress_Vendor

	con := db.CreateCon()

	sqlStatement := "UPDATE laporan_vendor SET status_laporan=? WHERE id_laporan_vendor=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_laporan_vendor)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	var id_jdl []string
	var id string

	sqlStatement = "SELECT id_kontrak FROM detail_laporan_vendor WHERE id_laporan_vendor=?"

	rows, err := con.Query(sqlStatement, id_laporan_vendor)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return res, err
		}
		id_jdl = append(id_jdl, id)
	}

	for i := 0; i < len(id_jdl); i++ {
		sqlstatemen_jdl := "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

		_ = con.QueryRow(sqlstatemen_jdl, id[i]).Scan(&RP.Id_kontrak, &RP.Working_Progess, &RP.Durasi, &RP.Working_Complate)

		if RP.Working_Complate == 1 {
			RP.Working_Progess = RP.Durasi
		}

		sqlStatement = "UPDATE kontrak_vendor SET working_progress=? WHERE id_kontrak=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(RP.Working_Progess, id[i])
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Upload_Foto_Laporan (done)V
func Upload_Foto_laporan_vendor(id_laporan_vendor string, writer http.ResponseWriter, request *http.Request) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM foto_laporan_vendor ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_LPV_FT := id_laporan_vendor + "-LPV-FT-" + strconv.Itoa(nm_str)

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name : ", handler.Filename)
	fmt.Println("File Size : ", handler.Size)
	fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

	var tempFile *os.File
	path := ""

	if strings.Contains(handler.Filename, "jpg") {
		path = "uploads/foto_laporan_vendor/" + id_LPV_FT + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/foto_laporan_vendor/" + id_LPV_FT + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/foto_laporan_vendor/" + id_LPV_FT + ".png"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan_vendor/", "Read"+"*.png")
	}

	if err != nil {
		return res, err
	}

	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return res, err2
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return res, err
	}

	fmt.Println("Success!!")
	fmt.Println(tempFile.Name())
	tempFile.Close()

	err = os.Rename(tempFile.Name(), path)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fmt.Println("new path:", tempFile.Name())

	sqlstatement := "INSERT INTO foto_laporan_vendor(co, id_foto_laporan_vendor, id_laporan_vendor, path) VALUE (?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nm_str, id_LPV_FT, id_laporan_vendor, path)

	if err != nil {
		return res, err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read_Foto_Laporan (done)V
func Read_Foto_Laporan_Vendor(id_laporan_vendor string) (tools2.Response, error) {
	var res tools2.Response
	var arr_Foto []vendor_all.Foto
	var Foto vendor_all.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT id_foto_laporan_vendor, id_laporan_vendor,path FROM foto_laporan_vendor WHERE foto_laporan_vendor.id_laporan_vendor=? ORDER BY co DESC "

	rows, err := con.Query(sqlStatement, id_laporan_vendor)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Foto.Id_Foto_Laporan_Vendor, &Foto.Id_Laporan_Vendor, &Foto.Path)

		if err != nil {
			return res, err
		}
		arr_Foto = append(arr_Foto, Foto)
	}

	if arr_Foto == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_Foto
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_Foto
	}

	return res, nil
}

//See-Task-Di-Input-Laporan (done)
func See_Task_Vendor(tanggal_laporan string, id_proyek string, id_laporan_vendor string) (tools2.Response, error) {
	var res tools2.Response

	var rt_lp vendor_all.Read_Task_Laporan_Vendor
	var arr_rt_lp []vendor_all.Read_Task_Laporan_Vendor

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	con := db.CreateCon()

	if id_laporan_vendor == "" {

		sqlStatement := "SELECT id_kontrak,nama_vendor,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),penkerjaan_vendor,working_progress FROM kontrak_vendor JOIN vendor v ON v.id_master_vendor=kontrak_vendor.id_MV WHERE tanggal_pengerjaan_dimulai<=? && tanggal_pengerjaan_berakhir>=? && working_progress != DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) && id_proyek=?"

		rows, err := con.Query(sqlStatement, date_sql, date_sql, id_proyek)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			durasi := 0

			err = rows.Scan(&rt_lp.Id_Kontrak, &rt_lp.Nama_Vendor, &durasi, &rt_lp.Pekerjaan_Vendor, &rt_lp.Progress)

			progress_float := float64(rt_lp.Progress)
			durasi_float := float64(durasi)

			rt_lp.Progress = int(math.Round((progress_float / durasi_float) * 100))

			if err != nil {
				return res, err
			}

			arr_rt_lp = append(arr_rt_lp, rt_lp)
		}

	} else {

		sqlStatement := "SELECT kontrak_vendor.id_kontrak,nama_vendor,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),penkerjaan_vendor,working_progress,ifnull(check_box,0) FROM kontrak_vendor JOIN vendor v ON v.id_master_vendor=kontrak_vendor.id_MV LEFT JOIN detail_laporan_vendor dlv on kontrak_vendor.id_kontrak = dlv.id_kontrak WHERE (tanggal_pengerjaan_dimulai<=? && tanggal_pengerjaan_berakhir>=? && working_progress != DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai) && id_proyek=?) || id_laporan_vendor=?"

		rows, err := con.Query(sqlStatement, date_sql, date_sql, id_proyek, id_laporan_vendor)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			durasi := 0

			err = rows.Scan(&rt_lp.Id_Kontrak, &rt_lp.Nama_Vendor, &durasi, &rt_lp.Pekerjaan_Vendor, &rt_lp.Progress, &rt_lp.Check_Box)

			progress_float := float64(rt_lp.Progress)
			durasi_float := float64(durasi)

			rt_lp.Progress = int(math.Round((progress_float / durasi_float) * 100))

			if err != nil {
				return res, err
			}

			arr_rt_lp = append(arr_rt_lp, rt_lp)
		}

	}

	if arr_rt_lp == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_rt_lp
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_rt_lp
	}

	return res, nil
}

//delete laporan vendor (done)
func Delete_laporan_Vendor(id_laporan_vendor string) (tools2.Response, error) {
	var res tools2.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan_vendor WHERE id_laporan_vendor=?"

	_ = con.QueryRow(sqlStatement, id_laporan_vendor).Scan(&st.Status)

	if st.Status == 0 {

		id := ""

		Sql := "SELECT id_detail_laporan_vendor FROM detail_laporan_vendor WHERE id_laporan_vendor=?"

		_ = con.QueryRow(Sql, id_laporan_vendor).Scan(&id)

		if id != "" {

			var read_dt_lp vendor_all.Detail_Laporan_Vendor_Update
			var arr_read_dt_lp []vendor_all.Detail_Laporan_Vendor_Update
			var rp vendor_all.Progress_Vendor

			sqlStatement := "SELECT co, id_detail_laporan_vendor,id_kontrak,check_box FROM detail_laporan_vendor WHERE id_laporan_vendor=?"

			rows, err := con.Query(sqlStatement, id_laporan_vendor)

			defer rows.Close()

			if err != nil {
				return res, err
			}

			for rows.Next() {
				err = rows.Scan(&read_dt_lp.Co, &read_dt_lp.Id_Detail_Laporan_Vendor, &read_dt_lp.Id_Kontrak_Vendor, &read_dt_lp.Check_Box)

				if err != nil {
					return res, err
				}
				arr_read_dt_lp = append(arr_read_dt_lp, read_dt_lp)
			}

			fmt.Println(arr_read_dt_lp, len(arr_read_dt_lp))

			for i := 0; i < len(arr_read_dt_lp); i++ {
				//Update Penjadwalan
				fmt.Println("Masuk")

				sqlStatement = "SELECT id_kontrak,working_progress,DATEDIFF(tanggal_pengerjaan_berakhir,tanggal_pengerjaan_dimulai),working_complate FROM kontrak_vendor WHERE id_kontrak=?"

				_ = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Id_Kontrak_Vendor).Scan(&rp.Id_kontrak, &rp.Working_Progess, &rp.Durasi, &rp.Working_Complate)

				if rp.Working_Complate == 1 && arr_read_dt_lp[i].Check_Box == 1 {
					rp.Working_Complate = 0
					rp.Working_Progess--
				} else if rp.Working_Complate == 1 {
					rp.Working_Complate = 0
				} else if rp.Working_Complate == 0 {
					rp.Working_Progess--
				}

				sqlStatement = "UPDATE kontrak_vendor SET working_progress=?,working_complate=? WHERE id_kontrak=?"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(rp.Working_Progess, rp.Working_Complate, rp.Id_kontrak)

				if err != nil {
					return res, err
				}

				//update penjadwalan
				var ARR_Id_Detail_Laporan []vendor_all.Id_Detail_Laporan_Vendor
				var Id_Detail_Laporan vendor_all.Id_Detail_Laporan_Vendor

				progres_lama := 0

				sqlStatement = "SELECT COUNT(id_detail_laporan_vendor) FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co < ? && id_kontrak=?"

				err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor).Scan(&progres_lama)

				sqlStatement = "SELECT id_detail_laporan_vendor FROM detail_laporan_vendor JOIN laporan_vendor lv on lv.id_laporan_vendor = detail_laporan_vendor.id_laporan_vendor WHERE lv.co > ? && id_kontrak=? ORDER BY lv.co ASC"

				rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Kontrak_Vendor)

				defer rows.Close()

				for rows.Next() {
					err = rows.Scan(&Id_Detail_Laporan.Id_Detail_Laporan_Vendor)

					if err != nil {
						return res, err
					}
					ARR_Id_Detail_Laporan = append(ARR_Id_Detail_Laporan, Id_Detail_Laporan)
				}

				for x := 0; x < len(ARR_Id_Detail_Laporan); x++ {
					progres_lama++
					progress_float := float64(progres_lama)
					durasi_float := float64(rp.Durasi)

					progress := int(math.Round((progress_float / durasi_float) * 100))

					sqlStatement = "UPDATE detail_laporan SET progress=? WHERE id_detail_laporan=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_Detail_Laporan_Vendor)

					if err != nil {
						return res, err
					}
				}

				sqlstatement := "DELETE FROM detail_laporan_vendor WHERE id_laporan_vendor=?"

				fmt.Println(sqlstatement)

				stmt, err = con.Prepare(sqlstatement)

				if err != nil {
					return res, err
				}

				result, err := stmt.Exec(id_laporan_vendor)

				if err != nil {
					return res, err
				}

				_, err = result.RowsAffected()

				if err != nil {
					return res, err
				}
			}

		}

		sqlStatement := "SELECT path FROM foto_laporan_vendor WHERE id_laporan_vendor=? "

		rows, err := con.Query(sqlStatement, id_laporan_vendor)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			path := ""
			err = rows.Scan(&path)
			path = "./" + path
			_ = os.Remove(path)
		}

		sqlstatement := "DELETE FROM foto_laporan_vendor WHERE id_laporan_vendor=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_laporan_vendor)

		if err != nil {
			return res, err
		}

		sqlstatement = "DELETE FROM laporan_vendor WHERE id_laporan_vendor=?"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err = stmt.Exec(id_laporan_vendor)

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
		res.Message = "Suksess"
	}

	return res, nil
}

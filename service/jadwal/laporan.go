package jadwal

import (
	"Skripsi/config/db"
	"Skripsi/models/jadwal"
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

//Input-Laporan (done)V
func Input_Laporan(id_proyek string, laporan string, tanggal_laporan string, id_penjadwalan string, check string) (tools2.Response, error) {
	var res tools2.Response
	var RP jadwal.Progress

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM laporan ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_laporan := "LP-" + strconv.Itoa(nm_str)

	sqlStatement := "INSERT INTO laporan (co, id_laporan, id_proyek, laporan, tanggal_laporan, foto_laporan, status_laporan) values(?,?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nm_str, id_laporan, id_proyek, laporan, date_sql, "", 0)

	if err != nil {
		return res, err
	}

	ip := tools2.String_Separator_To_String(id_penjadwalan)
	ck := tools2.String_Separator_To_Int(check)

	//masih salah kurang pengecekan udah finish ta blm e
	for i := 0; i < len(ip); i++ {
		progress := 0
		if ck[i] == 1 {
			sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			err = con.QueryRow(sqlstatemen_jdl, ip[i]).Scan(&RP.Id_penjadwalan, &RP.Progress,
				&RP.Durasi, &RP.Complate)

			if RP.Durasi == RP.Progress+1 {
				RP.Progress++
				RP.Complate = 1
			} else {
				RP.Progress = RP.Progress + 1
			}

			sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(RP.Progress, RP.Complate, ip[i])

			progress_float := float64(RP.Progress)
			durasi_float := float64(RP.Durasi)

			progress = int(math.Round((progress_float / durasi_float) * 100))

		} else if ck[i] == 2 {

			sqlStatement = "UPDATE penjadwalan SET complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(1, ip[i])

			progress = 100
		}

		nm_str_DLP := 0

		Sqlstatement := "SELECT co FROM detail_laporan ORDER BY co DESC Limit 1"

		_ = con.QueryRow(Sqlstatement).Scan(&nm_str_DLP)

		nm_str_DLP = nm_str_DLP + 1

		id_detail_laporan := "DLP-" + strconv.Itoa(nm_str_DLP)

		sqlStatement := "INSERT INTO detail_laporan (co, id_detail_laporan, id_laporan, id_jadwal, checkbox,progress) values(?,?,?,?,?,?)"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(nm_str_DLP, id_detail_laporan, id_laporan, ip[i], ck[i], progress)
	}

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

//Read-Laporan (done)V
func Read_Laporan(id_Proyek string) (tools2.Response, error) {
	var res tools2.Response

	var arr_lp []jadwal.Read_Laporan
	var lp jadwal.Read_Laporan

	//Read Laporan (Header)
	con := db.CreateCon()

	sqlStatement := "SELECT id_laporan, laporan, DATE_FORMAT(tanggal_laporan, '%d-%m%-%Y'),status_laporan FROM laporan WHERE laporan.id_Proyek=? ORDER BY tanggal_laporan desc"

	rows, err := con.Query(sqlStatement, id_Proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&lp.Id_laporan, &lp.Laporan, &lp.Tanggal_laporan, &lp.Status_laporan)

		if err != nil {
			return res, err
		}
		arr_lp = append(arr_lp, lp)
	}

	for i := 0; i < len(arr_lp); i++ {
		//Read id penjadwalan
		var arr_id jadwal.Detail_Laporan_TB

		sqlStatement := "SELECT id_jadwal,progress FROM detail_laporan WHERE id_laporan=?"

		rows, err := con.Query(sqlStatement, arr_lp[i].Id_laporan)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			var dl jadwal.Detail_Laporan
			err = rows.Scan(&arr_id.Id_Penjadwalan, &dl.Progress)

			if err != nil {
				return res, err
			}

			durasi := 0
			complate := 0

			sqlstatemen_jdl := "SELECT id_penjadwalan,nama_task,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlstatemen_jdl, arr_id.Id_Penjadwalan).Scan(&dl.Id_Penjadwalan,
				&dl.Nama_Sub_Pekerjaan, &durasi, &complate)

			if complate == 1 && durasi != dl.Progress {
				dl.Progress = durasi
			}

			arr_lp[i].Detail_Laporan = append(arr_lp[i].Detail_Laporan, dl)

		}

	}

	if arr_lp == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_lp
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_lp
	}

	return res, nil
}

//Update-Laporan (done) (done)V
func Update_Laporan(id_laporan string, laporan string, id_penjadwalan string, check string) (tools2.Response, error) {

	var res tools2.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan WHERE id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&st.Status)

	if st.Status == 0 {

		var read_dt_lp jadwal.Detail_Laporan_Update
		var arr_read_dt_lp []jadwal.Detail_Laporan_Update
		var rp jadwal.Progress
		var RP jadwal.Progress

		//mengurus detail laporan lama yang tidak di pakai lagi
		id_br := tools2.String_Separator_To_String(id_penjadwalan)
		ck := tools2.String_Separator_To_Int(check)

		sqlStatement = "SELECT co,id_detail_laporan,id_jadwal,checkbox FROM detail_laporan"

		q1 := " WHERE id_jadwal NOT IN ("

		for i := 0; i < len(id_br); i++ {
			if i == len(id_br)-1 {
				q1 = q1 + "'" + id_br[i] + "') && id_laporan = " + "'" + id_laporan + "' "
			} else {
				q1 = q1 + "'" + id_br[i] + "' , "
			}
		}

		sqlStatement = sqlStatement + q1

		rows, err := con.Query(sqlStatement)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&read_dt_lp.Co, &read_dt_lp.Id_detail_laporan, &read_dt_lp.Id_Penjadwalan, &read_dt_lp.Check_Box)

			if err != nil {
				return res, err
			}
			arr_read_dt_lp = append(arr_read_dt_lp, read_dt_lp)
		}

		fmt.Println(arr_read_dt_lp, len(arr_read_dt_lp))

		for i := 0; i < len(arr_read_dt_lp); i++ {
			//Update Penjadwalan
			fmt.Println("masuk")

			sqlStatement = "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Id_Penjadwalan).Scan(&rp.Id_penjadwalan,
				&rp.Progress, &rp.Durasi, &rp.Complate)

			if rp.Complate == 1 && arr_read_dt_lp[i].Check_Box == 1 {
				rp.Complate = 0
				rp.Progress--
			} else if rp.Complate == 1 {
				rp.Complate = 0
			} else if rp.Complate == 0 {
				rp.Progress--
			}

			sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(rp.Progress, rp.Complate, rp.Id_penjadwalan)

			if err != nil {
				return res, err
			}

			//update penjadwalan
			var ARR_Id_Detail_Laporan []jadwal.Id_Detail_Laporan
			var Id_Detail_Laporan jadwal.Id_Detail_Laporan

			progres_lama := 0

			sqlStatement = "SELECT COUNT(id_detail_laporan) FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE laporan.co < ? && id_jadwal=?"

			err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan).Scan(&progres_lama)

			sqlStatement = "SELECT id_detail_laporan FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE l.co > ? && id_jadwal=? ORDER BY l.co ASC"

			rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan)

			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&Id_Detail_Laporan.Id_detail_laporan)

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

				_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_detail_laporan)

				if err != nil {
					return res, err
				}
			}

			// menghilangkan detail laporan yang gak perlu --

			q2 := " WHERE id_detail_laporan IN ("

			for j := 0; j < len(arr_read_dt_lp); j++ {
				if j == len(arr_read_dt_lp)-1 {
					q2 = q2 + "'" + arr_read_dt_lp[i].Id_detail_laporan + "') && id_laporan = " + "'" + id_laporan + "' "
				} else {
					q2 = q2 + "'" + arr_read_dt_lp[i].Id_detail_laporan + "' , "
				}
			}

			sqlstatement := "DELETE FROM detail_laporan"

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

		//mengurus detail laporan baru
		for i := 0; i < len(id_br); i++ {

			fmt.Println("Masuk 2")

			id_detail_laporan := ""
			check_box := 0

			sqlStatement = "SELECT id_detail_laporan,checkbox FROM detail_laporan WHERE id_laporan=? && id_jadwal=?"

			_ = con.QueryRow(sqlStatement, id_laporan, id_br[i]).Scan(&id_detail_laporan, &check_box)

			fmt.Println(id_detail_laporan)

			if id_detail_laporan == "" {
				progress := 0
				if ck[i] == 1 {
					fmt.Println(id_br[i])
					sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

					err = con.QueryRow(sqlstatemen_jdl, id_br[i]).Scan(&RP.Id_penjadwalan, &RP.Progress, &RP.Durasi, &RP.Complate)

					if err != nil {
						return res, err
					}

					if RP.Durasi == RP.Progress+1 {
						RP.Progress++
						RP.Complate = 1
					} else {
						RP.Progress = RP.Progress + 1
					}

					sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(RP.Progress, RP.Complate, id_br[i])

					if err != nil {
						return res, err
					}

					var ARR_Id_Detail_Laporan []jadwal.Id_Detail_Laporan
					var Id_Detail_Laporan jadwal.Id_Detail_Laporan

					progres_lama := 0

					sqlStatement = "SELECT COUNT(id_detail_laporan) FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE laporan.co < ? && id_jadwal=?"

					err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan).Scan(&progres_lama)

					progres_lama = progres_lama + 1

					progress_float := float64(progres_lama)
					durasi_float := float64(RP.Durasi)

					progress = int(math.Round((progress_float / durasi_float) * 100))

					sqlStatement = "SELECT id_detail_laporan FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE l.co > ? && id_jadwal=? ORDER BY l.co ASC"

					rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan)

					defer rows.Close()

					for rows.Next() {
						err = rows.Scan(&Id_Detail_Laporan.Id_detail_laporan)

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

						_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_detail_laporan)

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

					_, err = stmt.Exec(1, id_br[i])

					if err != nil {
						return res, err
					}

					progress = 100
				}

				fmt.Println("MASUK 4")
				nm_str_DLP := 0

				Sqlstatement := "SELECT co FROM detail_laporan ORDER BY co DESC Limit 1"

				err = con.QueryRow(Sqlstatement).Scan(&nm_str_DLP)

				if err != nil {
					nm_str_DLP = 0
				}

				nm_str_DLP = nm_str_DLP + 1

				id_detail_laporan := "DLP-" + strconv.Itoa(nm_str_DLP)

				sqlStatement := "INSERT INTO detail_laporan (co, id_detail_laporan, id_laporan, id_jadwal, checkbox,progress) values(?,?,?,?,?,?)"

				stmt, err := con.Prepare(sqlStatement)

				if err != nil {
					return res, err
				}

				_, err = stmt.Exec(nm_str_DLP, id_detail_laporan, id_laporan, id_br[i], ck[i], progress)

				if err != nil {
					return res, err
				}

			} else if check_box != ck[i] {
				progress := 0
				if ck[i] == 2 {

					sqlStatement = "UPDATE penjadwalan SET complate=? WHERE id_penjadwalan=?"

					stmt, err := con.Prepare(sqlStatement)

					if err != nil {
						return res, err
					}

					_, err = stmt.Exec(1, id_br[i])

					progress = 100

				} else if ck[i] == 1 {
					prg_lama := 0
					durasi := 0
					sqlStatement = "SELECT durasi,progress FROM penjadwalan WHERE id_penjadwalan=?"

					err = con.QueryRow(sqlStatement, id_br[i]).Scan(&durasi, &prg_lama)

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

					_, err = stmt.Exec(cplt, prg_lama, id_br[i])

					progress_float := float64(prg_lama)
					durasi_float := float64(durasi)

					progress = int(math.Round((progress_float / durasi_float) * 100))
				}

				sqlStatement = "UPDATE detail_laporan SET progress=? WHERE id_detail_laporan=?"

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

		sqlStatement = "UPDATE laporan SET laporan=? WHERE id_laporan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(laporan, id_laporan)

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

//Update-Status-Laporan (done)V
func Update_Status_Laporan(id_laporan string) (tools2.Response, error) {
	var res tools2.Response
	var RP jadwal.Progress

	con := db.CreateCon()

	sqlStatement := "UPDATE laporan SET status_laporan=? WHERE id_laporan=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(1, id_laporan)

	if err != nil {
		return res, err
	}

	rowschanged, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	var id_jdl []string
	var id string

	sqlStatement = "SELECT id_jadwal FROM detail_laporan WHERE id_laporan=?"

	rows, err := con.Query(sqlStatement, id_laporan)

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
		sqlstatemen_jdl := "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

		_ = con.QueryRow(sqlstatemen_jdl, id_jdl[i]).Scan(&RP.Id_penjadwalan, &RP.Progress,
			&RP.Durasi, &RP.Complate)

		if RP.Complate == 1 {
			RP.Progress = RP.Durasi
		}

		sqlStatement = "UPDATE penjadwalan SET progress=? WHERE id_penjadwalan=?"

		stmt, err := con.Prepare(sqlStatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(RP.Progress, id_jdl[i])
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"rows": rowschanged,
	}

	return res, nil
}

//Delete-laporan (done)V (blm_di_cb)
func Delete_Laporan(id_laporan string) (tools2.Response, error) {
	var res tools2.Response
	var st jadwal.Status_laporan

	con := db.CreateCon()

	sqlStatement := "SELECT status_laporan FROM laporan WHERE id_laporan=?"

	_ = con.QueryRow(sqlStatement, id_laporan).Scan(&st.Status)
	if st.Status == 0 {

		var read_dt_lp jadwal.Detail_Laporan_Update
		var arr_read_dt_lp []jadwal.Detail_Laporan_Update
		var rp jadwal.Progress

		//mengurus detail laporan lama yang tidak di pakai lagi

		sqlStatement = "SELECT co,id_detail_laporan,id_jadwal,checkbox FROM detail_laporan WHERE id_laporan=?"

		rows, err := con.Query(sqlStatement, id_laporan)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&read_dt_lp.Co, &read_dt_lp.Id_detail_laporan, &read_dt_lp.Id_Penjadwalan, &read_dt_lp.Check_Box)

			if err != nil {
				return res, err
			}
			arr_read_dt_lp = append(arr_read_dt_lp, read_dt_lp)
		}

		fmt.Println(arr_read_dt_lp, len(arr_read_dt_lp))

		for i := 0; i < len(arr_read_dt_lp); i++ {
			fmt.Println("masuk")

			sqlStatement = "SELECT id_penjadwalan,progress,durasi,complate FROM penjadwalan WHERE id_penjadwalan=?"

			_ = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Id_Penjadwalan).Scan(&rp.Id_penjadwalan,
				&rp.Progress, &rp.Durasi, &rp.Complate)

			if rp.Complate == 1 && arr_read_dt_lp[i].Check_Box == 1 {
				rp.Complate = 0
				rp.Progress--
			} else if rp.Complate == 1 {
				rp.Complate = 0
			} else if rp.Complate == 0 {
				rp.Progress--
			}

			sqlStatement = "UPDATE penjadwalan SET progress=?,complate=? WHERE id_penjadwalan=?"

			stmt, err := con.Prepare(sqlStatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(rp.Progress, rp.Complate, rp.Id_penjadwalan)

			if err != nil {
				return res, err
			}

			//update penjadwalan
			var ARR_Id_Detail_Laporan []jadwal.Id_Detail_Laporan
			var Id_Detail_Laporan jadwal.Id_Detail_Laporan

			progres_lama := 0

			sqlStatement = "SELECT COUNT(id_detail_laporan) FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE laporan.co < ? && id_jadwal=?"

			err = con.QueryRow(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan).Scan(&progres_lama)

			sqlStatement = "SELECT id_detail_laporan FROM detail_laporan JOIN laporan JOIN laporan l on detail_laporan.id_laporan = l.id_laporan WHERE l.co > ? && id_jadwal=? ORDER BY l.co ASC"

			rows, err = con.Query(sqlStatement, arr_read_dt_lp[i].Co, arr_read_dt_lp[i].Id_Penjadwalan)

			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&Id_Detail_Laporan.Id_detail_laporan)

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

				_, err = stmt.Exec(progress, ARR_Id_Detail_Laporan[x].Id_detail_laporan)

				if err != nil {
					return res, err
				}
			}

			// menghilangkan detail laporan yang gak perlu --

			sqlstatement := "DELETE FROM detail_laporan WHERE id_laporan=?"

			fmt.Println("SQL:", sqlstatement)

			stmt, err = con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			result, err := stmt.Exec(id_laporan)

			if err != nil {
				return res, err
			}

			_, err = result.RowsAffected()

			if err != nil {
				return res, err
			}
		}

		sqlstatement := "DELETE FROM laporan WHERE id_laporan=?"

		stmt, err := con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		result, err := stmt.Exec(id_laporan)

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

//See-Task-Di-Input-Laporan (done)V
func See_Task(tanggal_laporan string, id_proyek string, id_laporan string) (tools2.Response, error) {
	var res tools2.Response

	var rt_lp jadwal.Read_Task_Laporan
	var arr_rt_lp []jadwal.Read_Task_Laporan

	date, _ := time.Parse("02-01-2006", tanggal_laporan)
	date_sql := date.Format("2006-01-02")

	con := db.CreateCon()

	if id_laporan == "" {
		sqlStatement := "SELECT id_penjadwalan,nama_task,durasi,progress FROM penjadwalan WHERE tanggal_dimulai<=? && tanggal_selesai>=? && penjadwalan.progress != penjadwalan.durasi && id_proyek=?"

		rows, err := con.Query(sqlStatement, date_sql, date_sql, id_proyek)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			durasi := 0

			err = rows.Scan(&rt_lp.Id_penjadwalan, &rt_lp.Nama_Task, &durasi, &rt_lp.Progress)

			rt_lp.Progress = (rt_lp.Progress / durasi) * 100

			if err != nil {
				return res, err
			}

			arr_rt_lp = append(arr_rt_lp, rt_lp)
		}
	} else {

		sqlStatement := "SELECT id_penjadwalan,nama_task,durasi,penjadwalan.progress,ifnull(checkbox,0) FROM penjadwalan Left Join detail_laporan dl on penjadwalan.id_penjadwalan = dl.id_jadwal WHERE (tanggal_dimulai<=? && tanggal_selesai>=? && penjadwalan.progress != penjadwalan.durasi && id_proyek=?) || id_laporan=?"

		rows, err := con.Query(sqlStatement, date_sql, date_sql, id_proyek, id_laporan)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			durasi := 0

			err = rows.Scan(&rt_lp.Id_penjadwalan, &rt_lp.Nama_Task, &durasi, &rt_lp.Progress, &rt_lp.Check_box)

			rt_lp.Progress = (rt_lp.Progress / durasi) * 100

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

//Upload-Foto-Laporan (done)V
func Upload_Foto_Laporan(id_laporan string, writer http.ResponseWriter, request *http.Request) (tools2.Response, error) {
	var res tools2.Response

	con := db.CreateCon()

	nm_str := 0

	Sqlstatement := "SELECT co FROM foto_laporan ORDER BY co DESC Limit 1"

	_ = con.QueryRow(Sqlstatement).Scan(&nm_str)

	nm_str = nm_str + 1

	id_LP_FT := id_laporan + "-LP-FT-" + strconv.Itoa(nm_str)

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
		path = "uploads/foto_laporan/" + id_LP_FT + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/foto_laporan/" + id_LP_FT + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/foto_laporan/" + id_LP_FT + ".png"
		tempFile, err = ioutil.TempFile("uploads/foto_laporan/", "Read"+"*.png")
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

	sqlstatement := "INSERT INTO foto_laporan (co, id_foto_laporan, id_laporan, path_foto) values(?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nm_str, id_LP_FT, id_laporan, path)

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

//Read-Foto-Laporan (done)V
func Read_Foto_Laporan(id_laporan string) (tools2.Response, error) {
	var res tools2.Response
	var Arr_Foto []jadwal.Foto
	var Foto jadwal.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT id_foto_laporan, id_laporan, path_foto FROM foto_laporan WHERE foto_laporan.id_laporan=? ORDER BY co desc"

	rows, err := con.Query(sqlStatement, id_laporan)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&Foto.Id_foto_laporan, &Foto.Id_laporan, &Foto.Path_foto)

		if err != nil {
			return res, err
		}
		Arr_Foto = append(Arr_Foto, Foto)
	}

	if Arr_Foto == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = Arr_Foto
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = Arr_Foto
	}

	return res, nil
}

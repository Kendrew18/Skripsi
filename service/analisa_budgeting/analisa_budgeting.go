package analisa_budgeting

import (
	"Skripsi/config/db"
	"Skripsi/models/budgeting"
	"Skripsi/service/tools"
	"fmt"
	"math"
	"net/http"
	"time"
)

func WeekStart_End(tanggal string) (string, string) {

	date, _ := time.Parse("02-01-2006", tanggal)

	fmt.Println(int(date.Weekday()))

	end := ""

	start := ""

	if int(date.Weekday()) == 0 {

		date_start := date.AddDate(0, 0, -6)

		start = date_start.Format("2006-01-02")

		end = date.Format("2006-01-02")

	} else if int(date.Weekday()) == 1 {

		date_end := date.AddDate(0, 0, 6)

		start = date.Format("2006-01-02")

		end = date_end.Format("2006-01-02")

	} else {

		minus := -1 * (int(date.Weekday()) - 1)
		plus := 7 - int(date.Weekday())

		date_start := date.AddDate(0, 0, minus)
		date_end := date.AddDate(0, 0, plus)

		start = date_start.Format("2006-01-02")

		end = date_end.Format("2006-01-02")

	}

	return start, end
}

func Analisa_Budgeting(Tanggal_sekarang string, id_proyek string) (tools.Response, error) {
	con := db.CreateCon()
	var res tools.Response

	var analisa_budgeting budgeting.Analisa_Budgeting

	var Biaya_Mingguan budgeting.Biaya_Mingguan

	var D_A_B budgeting.Detail_Analisa_Budgeting
	var arr_D_A_B []budgeting.Detail_Analisa_Budgeting

	analisa_budgeting.Tanggal_Awal, analisa_budgeting.Tanggal_Akhir = WeekStart_End(Tanggal_sekarang)

	sqlStatement := "SELECT dp.sub_total,tanggal_dimulai, tanggal_selesai FROM penjadwalan JOIN detail_penawaran dp on penjadwalan.id_penawaran = dp.id_penawaran WHERE tanggal_dimulai<=? && penjadwalan.tanggal_selesai>=?"

	rows, err := con.Query(sqlStatement, analisa_budgeting.Tanggal_Akhir, analisa_budgeting.Tanggal_Awal)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	date_awal, _ := time.Parse("2006-01-02", analisa_budgeting.Tanggal_Awal)

	date_akhir, _ := time.Parse("2006-01-02", analisa_budgeting.Tanggal_Akhir)

	tot := int64(0)

	for rows.Next() {
		err = rows.Scan(&Biaya_Mingguan.Biaya_Mingguan, &Biaya_Mingguan.Tanggal_Awal, &Biaya_Mingguan.Tanggal_Akhir)

		if err != nil {
			return res, err
		}

		date_awal_jad, _ := time.Parse("2006-01-02", Biaya_Mingguan.Tanggal_Awal)

		date_akhir_jad, _ := time.Parse("2006-01-02", Biaya_Mingguan.Tanggal_Akhir)

		dur := date_akhir_jad.Sub(date_awal_jad)

		dur_int := int64(dur.Hours()/24) + int64(1)

		biaya_float := float64(Biaya_Mingguan.Biaya_Mingguan)
		dur_float := float64(dur_int)

		biaya_per_hari := int64(math.Round((biaya_float / dur_float) * 100))

		biaya_fix := int64(0)

		if (date_awal.Before(date_awal_jad) && date_akhir.After(date_akhir_jad)) || (date_awal.Before(date_awal_jad) && (date_akhir == date_akhir_jad)) || ((date_awal == date_awal_jad) && date_akhir.After(date_akhir_jad)) {
			difference := date_akhir_jad.Sub(date_awal_jad)

			selisih := int64(difference.Hours()/24) + int64(1)

			biaya_fix = biaya_per_hari * selisih

			fmt.Println(selisih)

		} else if date_awal.After(date_awal_jad) && date_akhir.Before(date_akhir_jad) || ((date_awal == date_awal_jad) && date_akhir.Before(date_akhir_jad)) || (date_awal.After(date_awal_jad) && (date_akhir == date_akhir_jad)) || ((date_awal_jad == date_awal) && (date_akhir_jad == date_akhir)) {

			difference := date_akhir.Sub(date_awal)

			selisih := int64(difference.Hours()/24) + int64(1)

			biaya_fix = biaya_per_hari * selisih

			fmt.Println(selisih)

		} else if date_awal.After(date_awal_jad) && date_akhir.After(date_akhir_jad) || (date_awal.After(date_awal_jad) && (date_akhir_jad == date_awal)) {

			fmt.Println("Masuk 3")

			difference := date_akhir_jad.Sub(date_awal)

			selisih := int64(difference.Hours()/24) + int64(1)

			biaya_fix = biaya_per_hari * selisih

			fmt.Println(selisih)

		} else if date_awal.Before(date_awal_jad) && date_akhir.Before(date_akhir_jad) || (date_akhir.Before(date_akhir_jad) && (date_akhir_jad == date_akhir)) {

			fmt.Println("Masuk 4")

			difference := date_akhir.Sub(date_awal_jad)

			selisih := int64(difference.Hours()/24) + int64(1)

			biaya_fix = biaya_per_hari * selisih

			fmt.Println(selisih)

		}
		tot = tot + biaya_fix
	}

	analisa_budgeting.Biaya_Mingguan = tot

	sqlStatement = "SELECT id_penawaran,judul FROM penawaran WHERE id_proyek=?"

	rows, err = con.Query(sqlStatement, id_proyek)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	tot_EV := int64(0)
	tot_AC := int64(0)
	tot_PV := int64(0)

	for rows.Next() {
		var D_S_P budgeting.Detail_Sub_Pekerjaan
		var arr_D_S_P []budgeting.Detail_Sub_Pekerjaan

		err = rows.Scan(&D_A_B.Id_penawaran, &D_A_B.Nama_judul)

		if err != nil {
			return res, err
		}

		sqlStatement = "SELECT dp.id_sub_pekerjaan,dp.nama_sub_pekerjaan,durasi,progress,SUM(r.nominal_pembayaran),dp.sub_total FROM penjadwalan JOIN detail_penawaran dp on penjadwalan.id_penawaran = dp.id_penawaran JOIN realisasi r on dp.id_sub_pekerjaan = r.id_sub_pekerjaan WHERE dp.id_penawaran=?"

		rows2, err := con.Query(sqlStatement, D_A_B.Id_penawaran)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			dur := 0
			prog := 0

			err = rows.Scan(&D_S_P.Id_Sub_Pekerjaan, &D_S_P.Nama_Sub_Pekerjaan, &dur, &prog, &D_S_P.AC, &D_S_P.PV)

			if err != nil {
				return res, err
			}

			dur_f := float64(dur)
			prog_f := float64(prog)

			D_S_P.Progress = int64(math.Round((dur_f / prog_f) * 100))

			D_S_P.EV = D_S_P.AC * D_S_P.Progress

			tot_EV = tot_EV + D_S_P.EV
			tot_AC = tot_AC + D_S_P.AC
			tot_PV = tot_PV + D_S_P.PV

			arr_D_S_P = append(arr_D_S_P, D_S_P)
		}

		D_A_B.Detail_Sub_Pekerjaan = arr_D_S_P

		arr_D_A_B = append(arr_D_A_B, D_A_B)
	}

	analisa_budgeting.CV = tot_EV - tot_AC
	analisa_budgeting.SV = tot_EV - tot_PV

	if tot_AC != 0 {
		analisa_budgeting.CPI = tot_EV / tot_AC
	} else {
		analisa_budgeting.CPI = 0
	}

	analisa_budgeting.SPI = tot_EV / tot_PV

	analisa_budgeting.Detail_Analisa_Budgeting = arr_D_A_B

	if arr_D_A_B == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = analisa_budgeting
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = analisa_budgeting
	}

	return res, nil
}

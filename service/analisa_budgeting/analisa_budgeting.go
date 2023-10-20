package analisa_budgeting

import (
	"fmt"
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

/*func Analisa_Budgeting(Tanggal_sekarang string) (tools.Response, error) {
	con := db.CreateCon()
	var res tools.Response

	var analisa_budgeting budgeting.Analisa_Budgeting

	var Biaya_Mingguan budgeting.Biaya_Mingguan
	var arr_Biaya_Mingguan []budgeting.Biaya_Mingguan

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

		if (date_awal.Before(date_awal_jad) && date_akhir.After(date_akhir_jad)) || (date_awal.Before(date_awal_jad) && (date_akhir == date_akhir_jad)) || (date_awal == date_awal_jad) && date_akhir.After(date_akhir_jad)){

		}
	}

	if arr_analisa_budgeting == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_analisa_budgeting
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_analisa_budgeting
	}

	return res, nil
}*/

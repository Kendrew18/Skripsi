package budgeting

import (
	"Skripsi/service/tools"
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

func Analisa_Budgeting(Tanggal_sekarang string) (tools.Response, error) {
	var res tools.Response

	return res, nil
}

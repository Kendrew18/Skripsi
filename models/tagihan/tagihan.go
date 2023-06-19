package tagihan

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/tools"
)

//Generate Id Tagihan
func Generate_Id_Tagihan() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_tagihan FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_tagihan=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Penawaran(perihal string, tanggal_pembarian_kwitansi string,
	tanggal_pembayaran string) (tools.Response, error) {
	var res tools.Response

	return res, nil
}

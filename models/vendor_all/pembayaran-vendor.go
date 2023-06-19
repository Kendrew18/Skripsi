package vendor_all

import (
	"Skripsi/db"
	str "Skripsi/struct_all"
	"Skripsi/struct_all/vendor_all"
	"Skripsi/tools"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generate_Id_Pembayaran_Vendor() int {
	var obj str.Generate_Id

	con := db.CreateCon()

	sqlStatement := "SELECT id_pv FROM generate_id"

	_ = con.QueryRow(sqlStatement).Scan(&obj.Id)

	no := obj.Id
	no = no + 1

	sqlstatement := "UPDATE generate_id SET id_pv=?"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return -1
	}

	stmt.Exec(no)

	return no
}

func Input_Pembayaran_Vendor(id_kontrak string, nomor_invoice string,
	jumlah_pembayaran int64, tanggal_pembayaran string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	nm := Generate_Id_Pembayaran_Vendor()

	nm_str := strconv.Itoa(nm)

	id_PV := "PV-" + nm_str

	foto := "uploads/images.png"

	sqlStatement := "INSERT INTO pembayaran_vendor (id_PV,id_kontrak,nomor_invoice,jumlah_pembayaran,tanggal_pembayaran ,foto_invoice) values(?,?,?,?,?,?)"

	date, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_sql := date.Format("2006-01-02")

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id_PV, id_kontrak, nomor_invoice, jumlah_pembayaran, date_sql, foto)

	stmt.Close()

	res.Status = http.StatusOK
	res.Message = "Sukses"

	return res, nil
}

func Read_Pembayaran_Vendor(id_kontrak string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []vendor_all.Read_Pembayaran_Vendor
	var invent vendor_all.Read_Pembayaran_Vendor

	con := db.CreateCon()

	sqlStatement := "SELECT id_PV, nomor_invoice, jumlah_pembayaran, DATE_FORMAT(tanggal_pembayaran, '%d-%m%-%Y') FROM pembayaran_vendor WHERE id_kontrak=? ORDER BY co ASC "

	rows, err := con.Query(sqlStatement, id_kontrak)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&invent.Id_PV, &invent.Nomor_invoice, &invent.Jumlah_Pembayaran, &invent.Tanggal_Pembayaran)
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

func Upload_Invoice(id_PV string, writer http.ResponseWriter, request *http.Request) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()
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
		path = "uploads/" + id_PV + ".jpg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpg")
	}
	if strings.Contains(handler.Filename, "jpeg") {
		path = "uploads/" + id_PV + ".jpeg"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.jpeg")
	}
	if strings.Contains(handler.Filename, "png") {
		path = "uploads/" + id_PV + ".png"
		tempFile, err = ioutil.TempFile("uploads/", "Read"+"*.png")
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

	sqlstatement := "UPDATE pembayaran_vendor SET foto_invoice=? WHERE id_PV=?"
	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(path, id_PV)

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

func Read_Foto_Pembayaran_vendor(id_pembayaran_vendor string) (tools.Response, error) {
	var res tools.Response
	var arr_invent []str.Foto
	var invent str.Foto

	con := db.CreateCon()

	sqlStatement := "SELECT foto_invoice FROM pembayaran_vendor WHERE id_PV=? "

	err := con.QueryRow(sqlStatement, id_pembayaran_vendor).Scan(&invent.Path_foto)

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

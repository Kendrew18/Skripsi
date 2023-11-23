package main

import (
	"Skripsi/config/db"
	"Skripsi/models"
	"Skripsi/models/vendor_all"
	"Skripsi/routes"
	"fmt"
	"github.com/appleboy/go-fcm"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func Ticker() {
	ticker1 := time.Now()

	fmt.Println(ticker1)

	//fmt.Println(ticker1.Format("2006/01/02"))
}

func sendFCMNotification(token string, title string, message string) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up FCM client
	client, err := fcm.NewClient(os.Getenv("FCM_KEY"))
	if err != nil {
		return err
	}

	// Create the FCM message
	msg := &fcm.Message{
		RegistrationIDs: []string{token},
		Notification: &fcm.Notification{
			Title: title,
			Body:  message,
		},
		//To:   token,
		//Data: map[string]interface{}{"title": title, "message": message},
	}

	// Send the FCM message
	_, err = client.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func Pop_Up_Notif() {

	jakartatime, _ := time.LoadLocation("Asia/Jakarta")
	tim := time.Now().In(jakartatime)

	fmt.Println(tim)

	tm := tim.Format("2006-01-02")
	fmt.Println(tm)

	var user models.User
	var arr_user []models.User
	var invent vendor_all.Read_Notif_Pop_up

	con := db.CreateCon()

	sqlst := "SELECT kode_user, nama, token, status_user FROM user WHERE token != ? && status_user != ?"

	rows_user, err := con.Query(sqlst, "", 3)

	defer rows_user.Close()

	//fmt.Println(err)

	if err != nil {
		fmt.Println(err)
	}

	for rows_user.Next() {
		err := rows_user.Scan(&user.Id_user, &user.Nama_user, &user.Token_user, &user.Status_akun)

		if err != nil {
			fmt.Println(err)
		}

		arr_user = append(arr_user, user)
	}

	sqlStatement := "SELECT id_notif, DATE_FORMAT(tanggal, '%d-%m%-%Y'), pesan ,status_pop_up_1,status_pop_up_2 FROM notif WHERE tanggal<=? && (notif.status_pop_up_1=0 || notif.status_pop_up_2=0)"

	rows, err := con.Query(sqlStatement, tm)

	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		err = rows.Scan(&invent.Id_notif, &invent.Tanggal, &invent.Pesan, &invent.Status_1, &invent.Status_2)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(arr_user); i++ {
			if arr_user[i].Token_user != "" {

				sqlstatement := ""

				if arr_user[i].Status_akun == 1 && invent.Status_1 == 0 {

					err = sendFCMNotification(arr_user[i].Token_user, "Hello "+arr_user[i].Nama_user, invent.Pesan)
					if err != nil {

						fmt.Println("client 1: ", err)
					}

					sqlstatement = "UPDATE notif SET status_pop_up_1=? WHERE id_notif=?"
				} else if arr_user[i].Status_akun == 2 && invent.Status_2 == 0 {
					err = sendFCMNotification(arr_user[i].Token_user, "Hello "+arr_user[i].Nama_user, invent.Pesan)
					if err != nil {
						fmt.Println(err)
					}

					sqlstatement = "UPDATE notif SET status_pop_up_2=? WHERE id_notif=?"
				}

				stmt, _ := con.Prepare(sqlstatement)

				_, _ = stmt.Exec(1, invent.Id_notif)
			}
		}
	}
}

func main() {
	db.Init()
	e := routes.Init()
	//Ticker()

	jakartatime, _ := time.LoadLocation("Asia/Jakarta")

	s := gocron.NewScheduler(jakartatime)

	s.WaitForScheduleAll()
	s.Every(5).Second().Do(Pop_Up_Notif)
	s.StartAsync()
	fmt.Println(s.IsRunning())

	e.Logger.Fatal(e.Start(":38600"))

}

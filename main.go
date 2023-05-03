package main

import (
	"Skripsi/db"
	"Skripsi/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}

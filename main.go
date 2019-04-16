package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/routes"
)

const (
	host     = "cryo.cdnm8viilrat.us-east-2.rds-preview.amazonaws.com"
	port     = 5432
	user     = "cryoadmin"
	password = "x3FhcrWDxnxCq9p"
	dbname   = "cryo"
)

func main() {
	err := db.Connect(host, port, user, password, dbname, false)
	defer db.Close()
	if err != nil {
		fmt.Println("[Cryo] Fatal error")
	} else {
		router := routes.Setup()

		fmt.Println("[Cryo] API listening on port 3333")
		log.Fatal(http.ListenAndServe(":3333", router))
	}
}

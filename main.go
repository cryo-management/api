package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cryo-management/api/db"
	"github.com/cryo-management/api/routes"
)

func main() {
	err := db.Connect()
	defer db.Close()

	if err != nil {
		fmt.Println("[Cryo] Fatal error")
	} else {
		router := routes.Setup()

		fmt.Println("[Cryo] API listening on port 3333")
		log.Fatal(http.ListenAndServe(":3333", router))
	}
}

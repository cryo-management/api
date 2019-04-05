package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/cryo-management/api/router"
	"github.com/cryo-management/api/db"
)

func main() {
	err := db.Connect()
	defer db.Close()
	
	if err != nil {
		fmt.Println("[Cryo] Fatal error")
	} else {
		router := router.Setup()

		fmt.Println("[Cryo] API listening on port 3333")
		log.Fatal(http.ListenAndServe(":3333", router))
	}	
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/config"
	"github.com/cryo-management/api/routes"
)

// Config is a global variable to the application
var Config config.Config

func main() {
	fmt.Println("[Cryo] Loading configuration file")
	Config, err := config.NewConfig("config.toml")
	if err != nil {
		fmt.Println("[Cryo] Error while trying to load configuration file")
	} else {
		fmt.Println("[Cryo] Configuration file loaded successfully")
		fmt.Println("[Cryo] Connecting to the database")
		err = db.Connect(Config.Host, Config.Port, Config.User, Config.Password, Config.DBName, false)
		defer db.Close()
		if err != nil {
			fmt.Println("[Cryo] Error while connecting to database")
		} else {
			fmt.Println("[Cryo] Database connected successfully")
			router := routes.Setup()

			fmt.Println("[Cryo] API listening on port 3333")
			log.Fatal(http.ListenAndServe(":3333", router))
		}
	}
}

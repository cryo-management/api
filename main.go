package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/andreluzz/go-sql-builder/db"
	"github.com/cryo-management/api/routes"
)

//Config defines the struct of system configs
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func main() {
	fmt.Println("[Cryo] Connecting to the database")
	var config Config
	toml.DecodeFile("config.toml", &config)
	err := db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
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

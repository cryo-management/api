package db

import (
	"database/sql"
	"errors"
	"fmt"

	// PostgresSQL library
	_ "github.com/lib/pq"
)

const (
	host     = "cryo.cdnm8viilrat.us-east-2.rds-preview.amazonaws.com"
	port     = 5432
	user     = "cryoadmin"
	password = "x3FhcrWDxnxCq9p"
	dbname   = "cryo"
)

//Database docs
type Database struct{}

//Query docs
func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if conn == nil {
		return nil, errors.New("[Cryo] Error: Database not connected")
	}
	return conn.Query(query, args...)
}

//Insert docs
func (db *Database) Insert(query string, args ...interface{}) (string, error) {
	var id string
	err := conn.QueryRow(query+" RETURNING id", args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

var conn *sql.DB

//Connect docs
func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	err := errors.New("")
	conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	fmt.Println("[Cryo] Database successfully connected")
	return nil

}

//Close docs
func Close() {
	conn.Close()
	fmt.Println("[Cryo] Database connection successfully closed")
}
